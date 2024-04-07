package src

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/caffeine-addictt/video-manager/src/utils"
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
		// Validate working directory exists and is writable
		dirPath, err := utils.ValidateDirectory(workingDir)
		if err != nil {
			fmt.Printf("Failed to validate working directory: %#v\n", workingDir)
			Debug(err.Error())
			os.Exit(1)
		}

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
		if len(argSet) == 0 {
			fmt.Println("No URL(s) were passed!")
			os.Exit(1)
		}

		// Handle downloading
		var waitGroup sync.WaitGroup
		waitGroup.Add(len(argSet))

		for url := range argSet {
			go func(url string) {
				defer waitGroup.Done()

				split := strings.Split(url, "/")
				downloadLocation := filepath.Clean(filepath.Join(dirPath, split[len(split)-1]))

				// Ensure file already does not exist
				Info("Checking if " + downloadLocation + " already exists")
				if _, err := os.Stat(downloadLocation); err == nil {
					fmt.Printf("File already exists for %s\n", downloadLocation)
					Debug("File: " + downloadLocation + " already exists for " + url)
					return
				}

				// Get File
				fmt.Printf("Downloading %s to %s", url, downloadLocation)
				Info("Getting url: " + url)

				request, err := http.NewRequest(http.MethodGet, url, http.NoBody)
				if err != nil {
					fmt.Println("Failed to create request: " + url)
					Debug(err.Error())
					return
				}

				response, err := http.DefaultClient.Do(request)
				if err != nil {
					fmt.Println("Failed to get url: " + url)
					Debug(err.Error())
					return
				}
				defer response.Body.Close()

				// Create file
				Info("Creating file at: " + downloadLocation)
				out, err := os.Create(downloadLocation)
				if err != nil {
					fmt.Println("Failed to create file at: " + downloadLocation)
					Debug(err.Error())
					return
				}
				defer out.Close()

				// Write to file
				Info("Writing " + url + " to " + downloadLocation)
				if _, err := io.Copy(out, response.Body); err != nil {
					fmt.Printf("Failed to write to file at: %s\n", downloadLocation)
					Debug(err.Error())
					return
				}

				fmt.Println("Downloaded " + url + " to " + downloadLocation)
			}(url)
		}

		waitGroup.Wait()
	},
}

func init() {
	rootCommand.AddCommand(getCommand)
	getCommand.Flags().StringVarP(&getFlags.inputFile, "file", "f", "", "Path to the input file containing the url(s)")
}
