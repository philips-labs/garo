package cmd_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/philips-labs/garo/cmd"
)

func TestParseDateUnknown(t *testing.T) {
	d := cmd.ParseDate("unknown")

	assert.WithinDuration(t, time.Now(), d, 1*time.Millisecond)
}

func TestParseDateRFC3339(t *testing.T) {
	exp := "2019-11-17T16:11:22Z"
	d := cmd.ParseDate(exp)

	assert.Equal(t, exp, d.Format(time.RFC3339))
}

func TestParseDateNonRFC3339(t *testing.T) {
	d := cmd.ParseDate("01 Jan 15 10:00 UTC")
	exp := time.Now()

	assert.Equal(t, exp.Format(time.RFC3339), d.Format(time.RFC3339))
}

func TestVersionCommands(t *testing.T) {
	version := cmd.VersionInfo{
		Version: "test",
		Commit:  "ab23f6",
		Date:    cmd.ParseDate("2019-11-17T16:11:22Z"),
	}
	vc := cmd.NewVersionCommander(version)
	rootCmd := &cobra.Command{
		Use:   "garo",
		Short: "Github Actions Runner Orchestrator Server",
	}
	vc.AddToCommand(rootCmd)

	exp := fmt.Sprintf(`%s

  version:  %s
  commit:   %s
  date:     %s

`, rootCmd.Short, version.Version, version.Commit, version.Date.Format(time.RFC3339))
	args := []string{"version"}

	for _, arg := range args {
		t.Run(arg, func(tt *testing.T) {
			output, err := executeCommand(rootCmd, arg)
			assert := assert.New(tt)
			assert.NoError(err)
			assert.Equal(exp, output)
		})
	}
}
