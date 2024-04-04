package src

import (
	"fmt"
	"os"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/spf13/cobra"
)

// Root command
var rootCommand = &cobra.Command{
	Use:   "video-manager",
	Short: "Video Manager CLI",
	Long: utils.Multiline(
		"Download, manage and view videos locally.",
		"Documentation is available at https://github.com/caffeine-addictt/video-manager",
	),
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
