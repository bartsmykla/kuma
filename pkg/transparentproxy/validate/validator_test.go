package validate_test

import (
	"context"
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kumahq/kuma/pkg/core"
	kuma_log "github.com/kumahq/kuma/pkg/log"
	tproxy_config "github.com/kumahq/kuma/pkg/transparentproxy/config"
	"github.com/kumahq/kuma/pkg/transparentproxy/validate"
)

var _ = Describe("Should Validate iptables rules", func() {
	l := core.NewLoggerTo(os.Stdout, kuma_log.InfoLevel).WithName("validator")

	Describe("generate default validator config", func() {
		It("ipv4", func() {
			// when
			validator := validate.NewValidator().IPv4(l)

			// then
			Expect(validator.GetServerListenIP()).To(Equal("127.0.0.1"))
		})

		It("ipv6", func() {
			// when
			validator := validate.NewValidator().IPv6(l, tproxy_config.IPFamilyModeDualStack)

			// then
			serverIP := validator.GetServerListenIP()
			Expect(serverIP).To(Equal("::1"))

			splitByCon := strings.Split(serverIP, ":")
			Expect(len(splitByCon)).To(BeNumerically(">", 2))
		})
	})

	It("should pass when connect to correct address", func() {
		// given
		ctx := context.Background()
		validator := validate.NewValidator().
			WithServerPort(0).
			WithClientConnectIP("127.0.0.1").
			IPv4(l)

		// when
		exitC := make(chan struct{})
		port, err := validator.RunServer(ctx, exitC)
		Expect(err).ToNot(HaveOccurred())
		err = validator.RunClient(ctx, port, exitC)

		// then
		Expect(err).ToNot(HaveOccurred())
	})

	It("should fail when no iptables rules setup", func() {
		// given
		ctx := context.Background()
		// just to make test faster and there should be no flakiness here because the connection will never establish successfully without the redirection
		validator := validate.NewValidator().
			WithServerPort(0).
			WithClientRetryInterval(30 * time.Millisecond).
			IPv4(l)

		// when
		exitC := make(chan struct{})
		_, err := validator.RunServer(ctx, exitC)
		Expect(err).ToNot(HaveOccurred())
		// by using 0, the client will generate a random port to connect, simulating the scenario in the real world
		err = validator.RunClient(ctx, 0, exitC)

		// then
		Expect(err).To(HaveOccurred())
		errMsg := err.Error()
		containsTimeout := strings.Contains(errMsg, "i/o timeout")
		containsRefused := strings.Contains(errMsg, "refused")
		Expect(containsTimeout || containsRefused).To(BeTrue())
	})
})
