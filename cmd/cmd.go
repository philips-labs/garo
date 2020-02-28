package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitConfig(cfgFile string, additionalConfigs func()) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return err
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".garo")
	}

	viper.SetEnvPrefix("garo")
	viper.BindEnv("gh_token")
	viper.AutomaticEnv()

	if additionalConfigs != nil {
		additionalConfigs()
	}

	return viper.ReadInConfig()
}

func SetDefaultAndFlagBinding(cmd *cobra.Command, key, flag string, value interface{}) {
	viper.SetDefault(key, value)
	viper.BindPFlag(key, cmd.PersistentFlags().Lookup(flag))
}

func IsConfigError(err error) bool {
	_, ok := err.(viper.ConfigFileNotFoundError)
	return ok
}
