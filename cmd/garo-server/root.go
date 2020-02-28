package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/philips-labs/garo/cmd"
	"github.com/philips-labs/garo/server"
)

var (
	listenAddr string
	cfgFile    string
)

var (
	versionCommander *cmd.VersionCommander
	rootCmd          = &cobra.Command{
		Use:   "garo-server",
		Short: "Github Actions Runner Orchestrator Server",
		Long: `Github Actions Runner Orchestrator allows you to manage
	your Github Selfhosted action runners. garo-server functions as a
	centralized store for garo-agent(s) to fetch the configurations and
	provides an API to manage the self hosted runners for your repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			if v, _ := cmd.Flags().GetBool("version"); v {
				cmd.Println(versionCommander.SprintVersion(cmd))
				return
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			logger, err := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
			if err != nil {
				log.Fatalf("Can't initialize zap logger: %v", err)
			}
			defer logger.Sync()

			srv := server.New(ctx, server.Config{
				Addr:   listenAddr,
				Logger: logger,
			})
			go func() {
				err := srv.Run()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("Failed to run the server", zap.Error(err))
				}
			}()

			logger.Info("Server is ready to handle requests", zap.String("addr", srv.Addr))

			srv.GracefulShutdown()
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.garo.yaml)")
	rootCmd.PersistentFlags().StringVar(&listenAddr, "listenAddr", ":8080", "server listen address")
	rootCmd.Flags().BoolP("version", "v", false, "shows version information")

	configCommander := &cmd.ConfigCommander{}
	configCommander.AddToCommand(rootCmd)

	initVersionCommander()
	versionCommander.AddToCommand(rootCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetEnvPrefix("garo")
		viper.BindEnv("gh_token")
		viper.SetConfigName(".garo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
