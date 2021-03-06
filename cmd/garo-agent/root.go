package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/philips-labs/garo/agent"
	"github.com/philips-labs/garo/cmd"
)

var (
	cfgFile string
)

var (
	versionCommander *cmd.VersionCommander
	rootCmd          = &cobra.Command{
		Use:   "garo-agent",
		Short: "Github Actions Runner Orchestrator Agent",
		Long: `Github Actions Runner Orchestrator allows you to manage
	your Github Selfhosted action runners. garo-agent applies the runner
	policies based on the configurations.`,
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

			cfg := agent.Config{
				ServerAddr:      viper.GetString("agent.serverAddress"),
				Logger:          logger,
				RefreshInterval: viper.GetDuration("agent.refreshInterval"),
				Repositories:    viper.GetStringSlice("agent.repositories"),
			}
			err = agent.Run(ctx, cfg)
			if err != nil {
				logger.Error("Failed to run Agent", zap.Error(err))
				return
			}

			logger.Info("Agent is running for", zap.Strings("repos", cfg.Repositories))

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			sig := <-quit

			logger.Info("Agent is shutting down", zap.String("reason", sig.String()))
		},
	}
)

func initConfig() {
	err := cmd.InitConfig(cfgFile, func() {
		cmd.SetDefaultAndFlagBinding(rootCmd, "agent.serverAddress", "server-addr", "http://localhost:8080")
		viper.SetDefault("agent.refreshInterval", 1*time.Second)
	})
	if err != nil && !cmd.IsConfigError(err) {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Using config file: ", viper.ConfigFileUsed())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.garo.yaml)")
	rootCmd.PersistentFlags().String("server-addr", "", "address to garo-server")

	rootCmd.Flags().BoolP("version", "v", false, "shows version information")

	configCommander := &cmd.ConfigCommander{}
	configCommander.AddToCommand(rootCmd)

	initVersionCommander()
	versionCommander.AddToCommand(rootCmd)
}
