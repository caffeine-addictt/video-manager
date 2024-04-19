package src

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Global variables
var (
	// Global
	home string

	// Flags
	cfgFile   utils.FilePathFlag
	cacheFile utils.FilePathFlag
	verbose   bool
	debug     bool

	// Working Directory
	workingDir utils.DirPathFlag
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

func init() {
	home_, err := homedir.Dir()
	if err != nil {
		fmt.Println("Failed to get home directory")
		Debug(err.Error())
		os.Exit(1)
	}

	home = home_
}

// Configuration
func init() {
	cobra.OnInitialize(initConfig)

	rootCommand.PersistentFlags().VarP(&cfgFile, "config", "c", "config file (default is $HOME/.video-manager)")
	if err := rootCommand.MarkPersistentFlagFilename("config"); err != nil {
		fmt.Println("Failed to mark flag -c as filename in root command")
		Debug(err.Error())
		os.Exit(1)
	}

	rootCommand.PersistentFlags().VarP(&cacheFile, "cache", "C", "cache file (default is $HOME/.video-manager_history)")
	if err := rootCommand.MarkPersistentFlagFilename("cache"); err != nil {
		fmt.Println("Failed to mark flag -C as filename in root command")
		Debug(err.Error())
		os.Exit(1)
	}
	viper.SetDefault("cache", filepath.Clean(filepath.Join(home, ".video-manager_history")))
	if err := viper.BindPFlag("cache", rootCommand.PersistentFlags().Lookup("cache")); err != nil {
		fmt.Println("Failed to bind persistent flag 'cache'")
		Debug(err.Error())
		os.Exit(1)
	}

	// Verbosity
	rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCommand.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug output")

	// Working directory
	rootCommand.PersistentFlags().VarP(&workingDir, "dir", "w", "Working directory (default is ~/Videos)")
	if err := rootCommand.MarkPersistentFlagDirname("dir"); err != nil {
		fmt.Println("Failed to mark flag -w as dirname in root command")
		Debug(err.Error())
		os.Exit(1)
	}
	viper.SetDefault("dir", "~/Videos")
	if err := viper.BindPFlag("dir", rootCommand.PersistentFlags().Lookup("dir")); err != nil {
		fmt.Println("Failed to bind persistent flag 'dir'")
		Debug(err.Error())
		os.Exit(1)
	}

	// Preferred extensions
	viper.SetDefault("preferred_extensions", []string{
		".mp4",
		".avi",
		".webm",
	})

	// Initialize subcommands
	initCache()
}

func initConfig() {
	viper.SetConfigType("yaml")

	if cfgFile != "" {
		// Reading provided configuration file
		viper.SetConfigFile(cfgFile.String())
		Debug("-c supplied at " + cfgFile.String())
	} else {
		// Reading configuration file from either pwd or $HOME
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".video-manager")
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

			if err := viper.SafeWriteConfigAs(filepath.Join(home, ".video-manager")); err != nil {
				fmt.Println("Failed to create configuration file at " + home + "/.video-manager")
				Debug(err.Error())
				os.Exit(1)
			}
			fmt.Println("Created default configuration file at " + home + "/.video-manager")
		} else {
			fmt.Println("Failed to read configuration file")
			Debug(err.Error())
			os.Exit(1)
		}
	}

	// Update cobra flags with viper environment
	viper.AutomaticEnv()
	rootCommand.Flags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) {
			if err := rootCommand.Flags().Set(f.Name, viper.GetString(f.Name)); err != nil {
				fmt.Printf("Failed to set flag '%s' to value '%s'\n", f.Name, viper.GetString(f.Name))
				os.Exit(1)
			}
		}
	})

	Debug("Loaded configuration from " + viper.ConfigFileUsed())

	// Making caching file `$HOME/.video-manager_history`
	Debug("Writing cache file if it does not exist at " + home + "/.video-manager_history")
	file, err := os.OpenFile(filepath.Clean(filepath.Join(home, ".video-manager_history")), os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		fmt.Println("Failed to write cache file at " + home + "/.video-manager_history")
		Debug(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	Debug("Ensured cache file exists at " + home + "/.video-manager_history")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
