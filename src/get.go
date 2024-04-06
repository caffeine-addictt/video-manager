package src

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var getFlags struct {
	inputFile string
}

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Get and download videos",
	Long:  `Get and download videos from passed file and url(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		// Turn all URLS to a map to eliminate duplicates
		// We map string: struct{} for the smallest memory footprint
		argSet := make(map[string]struct{})
		for _, arg := range args {
			argSet[arg] = struct{}{}
		}

		// Validate explicitly passed URL(s)
		if len(argSet) > 0 {
			Debug("Validating passed URL(s)")
			for rawURL := range argSet {
				if _, err := url.ParseRequestURI(rawURL); err != nil {
					fmt.Println("Invalid URL: " + rawURL)
					os.Exit(1)
				}
			}
		}

		// Fetch URLs from file
		if getFlags.inputFile != "" {
			Debug("-f was passed, reading url(s) from file")
			file, err := os.Open(getFlags.inputFile)
			if err != nil {
				fmt.Printf("Failed to read file at %s\n", getFlags.inputFile)
				Debug(err.Error())
				os.Exit(1)
			}
			Debug("Closing file at " + getFlags.inputFile)
			defer file.Close()

			// Read URLs from file, line by line
			preURLCount := len(argSet)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// Ignore duplicates
				if _, ok := argSet[scanner.Text()]; ok {
					Info("Skipping duplicate URL: " + scanner.Text())
					continue
				}

				// Validate URL
				if _, err := url.ParseRequestURI(scanner.Text()); err != nil {
					fmt.Println("Invalid URL: " + scanner.Text())
					os.Exit(1)
				}
				argSet[scanner.Text()] = struct{}{}
			}
			Info("Read " + fmt.Sprint(len(args)-preURLCount) + " url(s) from " + getFlags.inputFile)
		}

		// Ensure a URL was passed
		if len(argSet) <= 0 {
			fmt.Println("No URL(s) were passed!")
			os.Exit(1)
		}

		// TODO: Handle downloading with progress indication
	},
}

func init() {
	rootCommand.AddCommand(getCommand)
	getCommand.Flags().StringVarP(&getFlags.inputFile, "file", "f", "", "Path to the input file containing the url(s)")
}
