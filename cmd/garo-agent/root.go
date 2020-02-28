package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/philips-labs/garo/agent"
	"github.com/philips-labs/garo/cmd"
)

var (
	cfgFile    string
	serverAddr string
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

			err := agent.Run(ctx, agent.Config{
				ServerAddr: serverAddr,
			})
			if err != nil {
				panic(err)
			}
		},
	}
)

func initConfig() {
	err := cmd.InitConfig(cfgFile, func() {
		cmd.SetDefaultAndFlagBinding(rootCmd, "agent.server_address", "server-addr", "http://localhost:8080")
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
	rootCmd.PersistentFlags().StringVar(&serverAddr, "server-addr", "", "address to garo-server")

	rootCmd.Flags().BoolP("version", "v", false, "shows version information")

	configCommander := &cmd.ConfigCommander{}
	configCommander.AddToCommand(rootCmd)

	initVersionCommander()
	versionCommander.AddToCommand(rootCmd)
}
