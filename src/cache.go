package src

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/spf13/cobra"
)

// Commands
var cacheCommand = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
	Long:  "This is a group of commands for managing the cache",
}

var cacheListCommand = &cobra.Command{
	Use:     "list <pattern?>",
	Short:   "List cache contents",
	Aliases: []string{"ls"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Debug("Opening cache file at " + cacheFile.String())
		file, err := os.Open(cacheFile.String())
		if err != nil {
			fmt.Println("Failed to open cache file")
			Debug(err.Error())
			os.Exit(1)
		}
		defer file.Close()

		if len(args) == 0 {
			args = []string{""}
		}

		Info("Reading from cache file...")
		buffer := bufio.NewScanner(file)
		for buffer.Scan() {
			stripped := regexp.MustCompile(`^(https?://)?(www\.)?`).ReplaceAllString(buffer.Text(), "")

			if strings.HasPrefix(stripped, args[0]) {
				fmt.Println(buffer.Text())
			} else if matched, err := regexp.MatchString(args[0], buffer.Text()); err != nil && matched {
				fmt.Println(buffer.Text())
			}
		}
	},
}

var cacheRemoveCommand = &cobra.Command{
	Use:   "remove <pattern>",
	Short: "Remove cache entries matching the given pattern",
	Long: utils.Multiline(
		"Matches the given pattern in the cache file and removes the matching entries.",
		"First checks if it starts with the given pattern, then does a RegExp match if no match is found.",
	),
	SuggestFor: []string{"clear"},
	Aliases:    []string{"rm"},
	Args:       cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			fmt.Println(utils.Multiline(
				"No pattern provided!",
				"If you want to remove all entries, use the 'cache clear' command.",
				"",
				"Or run with -h|--help for usage.",
			))
			os.Exit(1)
		}

		Debug("Opening cache file at " + cacheFile.String())
		file, err := os.Open(cacheFile.String())
		if err != nil {
			fmt.Println("Failed to open cache file")
			Debug(err.Error())
			os.Exit(1)
		}

		// Make TEMP file
		Debug("Making temp file")
		tmp, err := os.CreateTemp("", "video-manager_*")
		if err != nil {
			fmt.Println("Failed to create temp file")
			Debug(err.Error())
			os.Exit(1)
		}
		defer os.Remove(tmp.Name())

		// Delete TEMP file on os exit
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		go func() {
			<-c
			Info("\nCaught interrupt, cleaning up...")
			if err := file.Close(); err != nil {
				fmt.Printf("\nFailed to close cache file: %s\n", err.Error())
			}
			if err := tmp.Close(); err != nil {
				fmt.Printf("\nFailed to close temp file: %s\n", err.Error())
			}
			if err := os.Remove(tmp.Name()); err != nil {
				fmt.Printf("\nFailed to remove temp file: %s\n", err.Error())
			}
			os.Exit(0)
		}()

		// Filter cache file
		Info("Filtering cache file...")
		readBuffer := bufio.NewScanner(file)
		writeBuffer := bufio.NewWriter(tmp)
		for readBuffer.Scan() {
			if !strings.HasPrefix(readBuffer.Text(), args[0]) {
				if _, err := writeBuffer.WriteString(readBuffer.Text() + "\n"); err != nil {
					fmt.Println("Failed to write to temp file")
					Debug(err.Error())
					os.Exit(1)
				}
				continue
			}

			if matched, err := regexp.MatchString(args[0], readBuffer.Text()); err == nil && !matched {
				if _, err := writeBuffer.WriteString(readBuffer.Text() + "\n"); err != nil {
					fmt.Println("Failed to write to temp file")
					Debug(err.Error())
					os.Exit(1)
				}
			}
		}

		// Write to TEMP
		Debug("Writing to temp file")
		if err := writeBuffer.Flush(); err != nil {
			fmt.Println("Failed to write temp file at " + cacheFile)
			Debug(err.Error())
			os.Exit(1)
		}

		// Close both files
		Debug("Closing cache and temp files")
		if err := file.Close(); err != nil {
			fmt.Println("Failed to close cache file")
			Debug(err.Error())
			os.Exit(1)
		}
		if err := tmp.Close(); err != nil {
			fmt.Println("Failed to close temp file")
			Debug(err.Error())
			os.Exit(1)
		}

		// Replace cache file with TEMP
		Info("Replacing cache file with filtered entries")
		if err := os.Rename(tmp.Name(), file.Name()); err != nil {
			fmt.Println("Failed to replace cache file at " + cacheFile)
			Debug(err.Error())
			os.Exit(1)
		}
	},
}

var cacheClearCommand = &cobra.Command{
	Use:     "clear",
	Short:   "Clear the cache",
	Aliases: []string{"wipe", "clr"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		Debug("Creating/Truncating cache file at " + cacheFile.String())
		file, err := os.OpenFile(cacheFile.String(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
		if err != nil {
			fmt.Println("Failed to clear the cache at " + cacheFile)
			Debug(err.Error())
			os.Exit(1)
		}
		defer file.Close()

		fmt.Println("Cleared cache at " + cacheFile)
	},
}

func initCacheCmd() {
	rootCommand.AddCommand(cacheCommand)
	cacheCommand.AddCommand(
		cacheListCommand,
		cacheClearCommand,
		cacheRemoveCommand,
	)
}
