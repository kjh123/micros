package upgrade

import (
	"fmt"
	"github.com/kjh123/micros/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the micros tools",
	Long:  "Upgrade the micros tools. Example: micros upgrade",
	Run:   Run,
}

// Run upgrade the micros tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall("github.com/kjh123/micros")
	if err != nil {
		fmt.Println(err)
	}
}
