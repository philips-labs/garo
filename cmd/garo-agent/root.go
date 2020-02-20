package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/philips-labs/garo/cmd"
)

var cfgFile string

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
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.garo.yaml)")

	rootCmd.Flags().BoolP("version", "v", false, "shows version information")
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
		viper.SetConfigName(".garo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
