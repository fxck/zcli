package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	loggerPackage "github.com/zerops-io/zcli/src/service/logger"
	paramsPackage "github.com/zerops-io/zcli/src/service/params"
	storagePackage "github.com/zerops-io/zcli/src/service/storage"
)

var (
	logger = loggerPackage.New(loggerPackage.Config{})
	params *paramsPackage.Handler
)

func ExecuteRootCmd(builtinToken string) {

	storage, err := createStorage()
	if err != nil {
		logger.Error(err)
		return
	}
	params = paramsPackage.New(logger, storage)

	rootCmd := &cobra.Command{
		Use: "zcli",
	}

	params.RegisterString(rootCmd, "restApiAddress", "https://app.zerops.dev", "address of rest api", paramsPackage.Persistent())
	params.RegisterString(rootCmd, "grpcApiAddress", "app.zerops.dev:20902", "address of grpc api", paramsPackage.Persistent())
	params.RegisterString(rootCmd, "vpnApiAddress", "vpn.app.zerops.dev", "address of vpn api", paramsPackage.Persistent())
	params.RegisterString(
		rootCmd, "token", builtinToken, "authentication token",
		paramsPackage.Persistent(),
		paramsPackage.FromTempData(func(data *storagePackage.Data) interface{} {
			return data.Token
		}),
	)

	rootCmd.AddCommand(deployCmd())
	rootCmd.AddCommand(startVpnCmd())
	rootCmd.AddCommand(stopVpnCmd())
	rootCmd.AddCommand(loginCmd())

	err = params.InitViper()
	if err != nil {
		logger.Error(err)
		return
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func regSignals(contextCancel func(), logger loggerPackage.Logger) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Info("\n", "signal:", sig)
		contextCancel()
	}()
}
