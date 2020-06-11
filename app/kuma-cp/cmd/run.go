package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	kuma_cmd "github.com/Kong/kuma/pkg/cmd"
	config_core "github.com/Kong/kuma/pkg/config/core"
	kuma_version "github.com/Kong/kuma/pkg/version"

	ui_server "github.com/Kong/kuma/app/kuma-ui/pkg/server"
	admin_server "github.com/Kong/kuma/pkg/admin-server"
	api_server "github.com/Kong/kuma/pkg/api-server"
	"github.com/Kong/kuma/pkg/config"
	kuma_cp "github.com/Kong/kuma/pkg/config/app/kuma-cp"
	"github.com/Kong/kuma/pkg/core"
	"github.com/Kong/kuma/pkg/core/bootstrap"
	mads_server "github.com/Kong/kuma/pkg/mads/server"
	sds_server "github.com/Kong/kuma/pkg/sds/server"
	xds_server "github.com/Kong/kuma/pkg/xds/server"
)

var (
	runLog = controlPlaneLog.WithName("run")
)

func newRunCmd() *cobra.Command {
	return newRunCmdWithOpts(runCmdOpts{
		SetupSignalHandler: core.SetupSignalHandler,
	})
}

type runCmdOpts struct {
	SetupSignalHandler func() (stopCh <-chan struct{})
}

func newRunCmdWithOpts(opts runCmdOpts) *cobra.Command {
	args := struct {
		configPath string
		kumaCpMode string
	}{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Launch Control Plane",
		Long:  `Launch Control Plane.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := config_core.ValidateCpMode(args.kumaCpMode); err != nil {
				return err
			}
			cfg := kuma_cp.DefaultConfig()
			err := config.Load(args.configPath, &cfg)
			if err != nil {
				runLog.Error(err, "could not load the configuration")
				return err
			}
			cfg.Mode = args.kumaCpMode
			rt, err := bootstrap.Bootstrap(cfg)
			if err != nil {
				runLog.Error(err, "unable to set up Control Plane runtime")
				return err
			}
			cfgForDisplay, err := config.ConfigForDisplay(&cfg)
			if err != nil {
				runLog.Error(err, "unable to prepare config for display")
				return err
			}
			cfgBytes, err := config.ToJson(cfgForDisplay)
			if err != nil {
				runLog.Error(err, "unable to convert config to json")
				return err
			}
			runLog.Info(fmt.Sprintf("Current config %s", cfgBytes))
			if cfg.Mode != config_core.Global {
				if err := sds_server.SetupServer(rt); err != nil {
					runLog.Error(err, "unable to set up SDS server")
					return err
				}
				if err := xds_server.SetupServer(rt); err != nil {
					runLog.Error(err, "unable to set up xDS server")
					return err
				}
				if err := api_server.SetupServer(rt); err != nil {
					runLog.Error(err, "unable to set up API server")
					return err
				}
				if err := mads_server.SetupServer(rt); err != nil {
					runLog.Error(err, "unable to set up Monitoring Assignment server")
					return err
				}
			}
			if err := admin_server.SetupServer(rt); err != nil {
				runLog.Error(err, "unable to set up Admin server")
				return err
			}
			if cfg.Mode != config_core.Local {
				if err := ui_server.SetupServer(rt); err != nil {
					runLog.Error(err, "unable to set up GUI server")
					return err
				}
			}

			runLog.Info("starting Control Plane", "version", kuma_version.Build.Version)
			if err := rt.Start(opts.SetupSignalHandler()); err != nil {
				runLog.Error(err, "problem running Control Plane")
				return err
			}

			runLog.Info("stopping Control Plane")
			return nil
		},
	}
	// flags
	cmd.PersistentFlags().StringVarP(&args.configPath, "config-file", "c", "", "configuration file")
	cmd.PersistentFlags().StringVar(&args.kumaCpMode, "mode", config_core.Standalone, kuma_cmd.UsageOptions("kuma cp modes", "standalone", "local", "global"))
	return cmd
}
