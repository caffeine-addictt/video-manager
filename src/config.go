package src

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Head
var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Manage your configuration",
}

// Sub
var configInitFlags struct {
	yes bool
}
var configInitCommand = &cobra.Command{
	Use:   "init <path?>",
	Short: "Initialize a new configuration file",
	Long: utils.Multiline(
		"Initialize a new configuration file at the specified path.",
		"If no path is specified, it will be created in the current directory.",
	),
	Aliases: []string{"create", "new", "reset"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var newConfigPath string
		if len(args) == 1 {
			newConfigPath = args[0]
		} else {
			newConfigPath = "./.video-manager.yml"
		}
		newConfigPath = filepath.Clean(newConfigPath)

		// Write config
		if err := viper.SafeWriteConfigAs(newConfigPath); err != nil {
			if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
				switch utils.InputPrompt("A configuration file already exists. Do you want to overwrite it? (y/n)") {
				case "y", "Y":
				default:
					os.Exit(1)
				}

				// Force write
				if err := viper.WriteConfigAs(newConfigPath); err != nil {
					fmt.Println("Failed to write configuration file")
					Debug(err.Error())
					os.Exit(1)
				}
			} else {
				fmt.Println("Failed to write configuration file")
				Debug(err.Error())
				os.Exit(1)
			}
		}
	},
}

var configWhereCommand = &cobra.Command{
	Use:   "which",
	Short: "Print the path to the loaded configuration file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.ConfigFileUsed())
	},
}

func initConfigCmd() {
	rootCommand.AddCommand(configCommand)
	configCommand.AddCommand(
		configInitCommand,
		configWhereCommand,
	)

	configInitCommand.Flags().BoolVarP(&configInitFlags.yes, "yes", "y", false, "Skip the confirmation prompt")
}
