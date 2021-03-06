package main

import (
	"github.com/kjh123/micros/internal/project"
	"github.com/kjh123/micros/internal/run"
	"github.com/kjh123/micros/internal/upgrade"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "micros",
	Version: "v1.0.2",
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
