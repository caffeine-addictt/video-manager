package src

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of video-manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.4")
	},
}

func init() {
	rootCommand.AddCommand(versionCommand)
}
