package cmd_test

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/philips-labs/garo/cmd"
	"github.com/philips-labs/garo/cmd/test"
)

var (
	token  string
	expCfg = func(token string) string { return "\nconfig:\n  gh_token:  " + token + "\n" }
)

func getRootCmdWithConfig() (*cobra.Command, *cmd.ConfigCommander) {
	cc := &cmd.ConfigCommander{}
	rootCmd := &cobra.Command{
		Use:   "garo",
		Short: "Github Actions Runner Orchestrator",
	}
	cc.AddToCommand(rootCmd)
	return rootCmd, cc
}

func initConfig(configPath string) error {
	viper.SetConfigName(".garo")
	viper.AddConfigPath(configPath)
	viper.SetEnvPrefix("garo")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	viper.BindEnv("gh_token")

	return viper.ReadInConfig()
}

func init() {
	wd, _ := os.Getwd()
	bytes := make([]byte, 10)
	rand.Read(bytes)
	token = fmt.Sprintf("%x", sha1.New().Sum(bytes))
	os.Setenv("GARO_GH_TOKEN", token)
	initConfig(filepath.Join(wd, ".."))
}

func TestSprintConfig(t *testing.T) {
	assert := assert.New(t)
	_, cc := getRootCmdWithConfig()
	cfg := cc.SprintConfig()
	assert.Equal(expCfg(token), cfg, "Expected configuration to match expectation")
}

func TestConfigCommand(t *testing.T) {
	assert := assert.New(t)
	rootCmd, _ := getRootCmdWithConfig()
	output, err := test.ExecuteCommand(rootCmd, "config")
	assert.NoError(err)
	assert.Equal(expCfg(token), output, "Expected configuration to be outputted in a different format")
}

func TestConfigPathRelativeToCwd(t *testing.T) {
	home, err := homedir.Dir()
	assert.NoError(t, err)

	_, cc := getRootCmdWithConfig()
	wd, err := os.Getwd()
	assert.NoError(t, err)

	testCases := []struct {
		key, input, exp string
	}{
		{key: "rel_wd_test_empty", input: "", exp: ""},
		{key: "rel_wd_test_absolute", input: "/var/lib/stuff", exp: "/var/lib/stuff"},
		{key: "rel_wd_test_relative", input: "./lib/stuff", exp: filepath.Join(wd, "lib", "stuff")},
		{key: "rel_wd_test_home", input: "~/stuff", exp: filepath.Join(home, "stuff")},
	}

	for _, tc := range testCases {
		t.Run(tc.key, func(tt *testing.T) {
			viper.Set(tc.key, tc.input)
			cc.ResolveConfigPathsRelativeToCwd(tc.key)
			assert.Equal(tt, tc.exp, viper.Get(tc.key))
		})
	}
}

func TestConfigPathRelativeToConfig(t *testing.T) {
	home, err := homedir.Dir()
	assert.NoError(t, err)

	wd, err := os.Getwd()
	assert.NoError(t, err)

	_, cc := getRootCmdWithConfig()

	testCases := []struct {
		key, input, exp string
	}{
		{key: "rel_config_test_empty", input: "", exp: ""},
		{key: "rel_config_test_absolute", input: "/var/lib/stuff", exp: "/var/lib/stuff"},
		{key: "rel_config_test_home", input: "~/stuff", exp: filepath.Join(home, "stuff")},
		{key: "rel_config_test_relative", input: "./lib/stuff", exp: filepath.Join(wd, "..", "lib", "stuff")},
	}

	for _, tc := range testCases {
		t.Run(tc.key, func(tt *testing.T) {
			viper.Set(tc.key, tc.input)
			cc.ResolveConfigPathsRelativeToConfig(tc.key)
			assert.Equal(tt, tc.exp, viper.Get(tc.key))
		})
	}
}
