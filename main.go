package main

import (
	"github.com/kjh123/micros/internal/project"
	"github.com/kjh123/micros/internal/proto"
	"github.com/kjh123/micros/internal/run"
	"github.com/kjh123/micros/internal/upgrade"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "micro",
	Short:   "micro: 本工具改自 kratos, 为适应此项目, 对代码有做改动, 如有问题参考原文档: https://go-kratos.dev/docs/getting-started/start.",
	Long:    `micro: 本工具改自 kratos, 为适应此项目, 对代码有做改动, 如有问题参考原文档: https://go-kratos.dev/docs/getting-started/start.`,
	Version: "v1.0.3",
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
