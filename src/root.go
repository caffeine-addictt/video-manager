package src

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Global variables
var (
	cfgFile string
	verbose bool
	debug   bool
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

// Configuration
func init() {
	cobra.OnInitialize(initConfig)
	rootCommand.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.video-manager)")

	// Verbosity
	rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCommand.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug output")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Failed to get home directory")
		Debug(err.Error())
		os.Exit(1)
	}

	if cfgFile != "" {
		// Reading provided configuration file
		viper.SetConfigFile(cfgFile)
		Debug("-c supplied at " + cfgFile)
	} else {
		// Reading configuration file from either pwd or $HOME
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".video-manager")
		viper.SetConfigType("yaml")
		Debug("-c not supplied, looking for configuration file in " + home + " or pwd")
	}

	// Read file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Configuration file not found")

			// Create configuration at $HOME if it does not exist
			if cfgFile != "" {
				os.Exit(1)
			}

			viper.SafeWriteConfigAs(filepath.Join(home, ".video-manager"))
			fmt.Println("Created default configuration file at " + home + "/.video-manager")
		} else {
			fmt.Println("Failed to read configuration file")
			Debug(err.Error())
			os.Exit(1)
		}
	}
	Info("Loaded configuration from " + viper.ConfigFileUsed())
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
