package validate

import (
	"context"
	"crypto/rand"
	std_errors "errors"
	"fmt"
	"math/big"
	"net"
	"net/netip"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-retry"

	"github.com/kumahq/kuma/pkg/core"
	tproxy_config "github.com/kumahq/kuma/pkg/transparentproxy/config"
	tproxy_consts "github.com/kumahq/kuma/pkg/transparentproxy/consts"
)

var log = core.Log.WithName("transparentproxy.validator")

const (
	serverListenPort      uint16 = 15006
	clientRetries                = 10
	clientRetriesInterval        = time.Second
)

type ValidatorBuilder struct {
	serverListenPort    uint16
	serverListenIP      netip.Addr
	clientConnectIP     netip.Addr
	clientRetryInterval time.Duration
}

func NewValidator() *ValidatorBuilder {
	return &ValidatorBuilder{
		clientRetryInterval: clientRetriesInterval,
		serverListenPort:    serverListenPort,
	}
}

func (vb *ValidatorBuilder) build(logName, serverListenIP, clientConnectIP string) *Validator {
	v := Validator{
		logger:              log.WithName(logName),
		serverListenPort:    vb.serverListenPort,
		serverListenIP:      netip.MustParseAddr(serverListenIP),
		clientConnectIP:     netip.MustParseAddr(clientConnectIP),
		clientRetryInterval: vb.clientRetryInterval,
	}

	if vb.clientConnectIP.IsValid() {
		v.clientConnectIP = vb.clientConnectIP
	}

	return &v
}

func (vb *ValidatorBuilder) IPv4() *Validator {
	return vb.build("ipv4", "127.0.0.1", "127.0.0.6")
}

func (vb *ValidatorBuilder) IPv6(mode tproxy_config.IPFamilyMode) *Validator {
	if mode == tproxy_config.IPFamilyModeIPv4 || mode == "" {
		return nil
	}

	if ok, _ := tproxy_config.HasLocalIPv6(); !ok  {
		return nil
	}

	return vb.build("ipv6", "::1", "::6")
}

func (vb *ValidatorBuilder) WithServerPort(port uint16) *ValidatorBuilder {
	vb.serverListenPort = port
	return vb
}

// Just for unit tests
func (vb *ValidatorBuilder) WithClientRetryInterval(interval time.Duration) *ValidatorBuilder {
	vb.clientRetryInterval = interval
	return vb
}

// Just for unit tests
func (vb *ValidatorBuilder) WithClientConnectIP(ip string) *ValidatorBuilder {
	vb.clientConnectIP = netip.MustParseAddr(ip)
	return vb
}

type Validator struct {
	logger              logr.Logger
	serverListenPort    uint16
	serverListenIP      netip.Addr
	clientConnectIP     netip.Addr
	clientRetryInterval time.Duration
}

// Just for unit tests
func (v *Validator) GetServerListenIP() string {
	return v.serverListenIP.String()
}

func (v *Validator) RunServer(
	ctx context.Context,
	exitC chan struct{},
) (uint16, error) {
	v.logger.Info("starting iptables validation")

	s := LocalServer{
		logger:   v.logger.WithValues("address", v.getServerAddress()),
		listenIP: []byte(v.serverListenIP.String()),
		address:  v.getServerAddress(),
	}

	readyC := make(chan struct{}, 1)
	errorC := make(chan error, 1)

	go func() {
		errorC <- s.Run(ctx, readyC, exitC)
	}()

	<-readyC

	select {
	case err := <-errorC:
		if err != nil {
			return 0, err
		}

		return 0, errors.New("server exited unexpectedly")
	default:
		return s.port, nil
	}
}

func (v *Validator) RunClient(
	ctx context.Context,
	serverPort uint16,
	exitC chan struct{},
) error {
	defer close(exitC)

	if err := retry.Do(
		ctx,
		retry.WithMaxRetries(clientRetries, retry.NewConstant(v.clientRetryInterval)), // backoff
		func(context.Context) error {
			if err := runLocalClient(v.clientConnectIP, serverPort); err != nil {
				v.logger.Info("failed to connect to the server, retrying...", "error", err)
				return retry.RetryableError(err)
			}
			v.logger.Info("client successfully connected to the server")
			return nil
		},
	); err != nil {
		return errors.Wrap(err, "client failed to connect to the verification server after retries")
	}

	v.logger.Info("validation successful, iptables rules applied correctly")

	return nil
}

func (v *Validator) Run(ctx context.Context) error {
	if v == nil {
		return nil
	}

	v.logger = v.logger.WithName(getLoggerName(v.serverListenIP))

	exitC := make(chan struct{})

	if _, err := v.RunServer(ctx, exitC); err != nil {
		return err
	}

	// by using 0, we make the client to generate a random port to connect verifying
	// the iptables rules are working
	return v.RunClient(ctx, 0, exitC)
}

func (v *Validator) getServerAddress() string {
	if v == nil {
		return ""
	}

	return net.JoinHostPort(
		v.serverListenIP.String(),
		fmt.Sprintf("%d", v.serverListenPort),
	)
}

func getLoggerName(addr netip.Addr) string {
	return strings.ToLower(tproxy_consts.IPTypeMap[addr.Is6()])
}

func RunValidation[T uint16 | tproxy_config.Port](
	ctx context.Context,
	serverPort T,
	mode tproxy_config.IPFamilyMode,
) error {
	return errors.Wrap(
		std_errors.Join(
			NewValidator().WithServerPort(uint16(serverPort)).IPv4().Run(ctx),
			NewValidator().WithServerPort(uint16(serverPort)).IPv6(mode).Run(ctx),
		),
		"validation failed",
	)
}

type LocalServer struct {
	logger   logr.Logger
	port     uint16
	address  string
	listenIP []byte
}

func (s *LocalServer) Run(
	ctx context.Context,
	readyC chan struct{},
	exitC chan struct{},
) error {
	s.logger.Info("server is listening")

	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", s.address)
	if err != nil {
		close(readyC)
		return err
	}

	s.port = uint16(l.Addr().(*net.TCPAddr).Port)

	go s.handleTcpConnections(l, exitC)
	close(readyC)
	<-exitC
	_ = l.Close()

	return nil
}

func (s *LocalServer) handleTcpConnections(l net.Listener, exitC chan struct{}) {
	for {
		conn, err := l.Accept()
		if err != nil {
			s.logger.Error(err, "failed to accept TCP connection")
			return
		}

		s.logger.Info("connection established")

		_, _ = conn.Write(s.listenIP)
		_ = conn.Close()

		select {
		case <-exitC:
			return
		default:
		}
	}
}

func runLocalClient(serverIP netip.Addr, serverPort uint16) error {
	var serverAddress string
	var laddr *net.TCPAddr
	var err error

	switch {
	case serverIP.Is6():
		if laddr, err = net.ResolveTCPAddr("tcp", "[::1]:0"); err != nil {
			return err
		}
	default:
		if laddr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:0"); err != nil {
			return err
		}
	}

	dialer := net.Dialer{LocalAddr: laddr, Timeout: 50 * time.Millisecond}

	// connections to all ports should be redirected to the server we support
	// a pre-configured port for testing purposes
	switch serverPort {
	case 0:
		randPort, _ := rand.Int(rand.Reader, big.NewInt(10000))
		serverAddress = net.JoinHostPort(
			serverIP.String(),
			fmt.Sprintf("%d", 20000+randPort.Int64()),
		)
	default:
		serverAddress = net.JoinHostPort(
			serverIP.String(),
			fmt.Sprintf("%d", serverPort),
		)
	}

	conn, err := dialer.Dial("tcp", serverAddress)
	// conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
