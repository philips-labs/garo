package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/philips-labs/garo/cmd/test"
)

func TestVersionOption(t *testing.T) {
	initVersionCommander()
	version := versionCommander.GetVersionInfo()
	exp := fmt.Sprintf(`%s

  version:  %s
  commit:   %s
  date:     %s

`, rootCmd.Short, version.Version, version.Commit, version.Date.Format(time.RFC3339))

	args := []string{"--version", "-v"}

	for _, arg := range args {
		t.Run(arg, func(tt *testing.T) {
			output, err := test.ExecuteCommand(rootCmd, arg)
			assert.NoError(tt, err)
			assert.Equal(tt, exp, output)
		})
	}
}
