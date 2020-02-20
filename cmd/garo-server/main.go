package main

import (
	"fmt"
	"os"

	"github.com/philips-labs/garo/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	v := cmd.VersionInfo{
		Version: version,
		Commit:  commit,
		Date:    cmd.ParseDate(date),
	}

	versionCommander = cmd.NewVersionCommander(v)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
