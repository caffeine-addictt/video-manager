package src

import (
	"bufio"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/spf13/cobra"
)

// Non URL similarity
// The resolved URLs from cache to the similarity score
type Similarity map[string]int8

// Command stuff
var getFlags struct {
	inputFile      utils.FilePathFlag
	strategy       utils.StrategyEnum
	maxConcurrency int
}

var getCommand = &cobra.Command{
	Use:   "get <url?>...",
	Short: "Get and download videos",
	Long:  `Get and download videos from passed file and url(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		// Warn on inefficient settings
		if getFlags.maxConcurrency == 1 && getFlags.strategy == utils.StrategyConcurrent {
			fmt.Println("WARNING: Setting -m to 1 with -s concurrent may not be efficient, please consider using -s synchronous instead.")
		}

		// Validate working directory exists and is writable
		dirPath, err := utils.ValidateDirectory(workingDir.String())
		if err != nil {
			fmt.Printf("Failed to validate working directory: %#v\n", workingDir)
			Debug(err.Error())
			os.Exit(1)
		}

		// Turn all URLS to a map to eliminate duplicates
		// We map string: struct{} for the smallest memory footprint
		argSet := make(map[string]struct{})
		nonURLSet := make(map[string]Similarity)

		for _, arg := range args {
			if _, err := url.ParseRequestURI(arg); err != nil {
				Debug("Invalid URL: " + arg)
				nonURLSet[arg] = make(Similarity, 10)
				continue
			}
			argSet[arg] = struct{}{}
		}

		// Fetch URLs from file
		if getFlags.inputFile != "" {
			Debug("-f was passed, reading url(s) from file")
			file, err := os.Open(getFlags.inputFile.String())
			if err != nil {
				fmt.Printf("Failed to read file at %s\n", getFlags.inputFile)
				Debug(err.Error())
				os.Exit(1)
			}
			Debug("Closing file at " + getFlags.inputFile.String())
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
					Debug("Invalid URL: " + scanner.Text())
					nonURLSet[scanner.Text()] = make(Similarity, 10)
					continue
				}
				argSet[scanner.Text()] = struct{}{}
			}
			Info("Read " + fmt.Sprint(len(args)-preURLCount) + " url(s) from " + getFlags.inputFile.String())
		}
		Debug("Valid URL(s): " + fmt.Sprint(argSet))
		Debug("Non URL(s): " + fmt.Sprint(nonURLSet))

		// Open cache file
		Debug("Opening cache file at " + cacheFile.String())
		file, err := os.OpenFile(cacheFile.String(), os.O_RDWR|os.O_APPEND, 0o600)
		if err != nil {
			fmt.Println("Failed to open cache file at " + cacheFile)
			Debug(err.Error())
			return
		}
		defer file.Close()

		// Provide completions from cache file
		if len(nonURLSet) > 0 {
			Info("Reading cache file to complete URL(s)...")
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// Skip if url already in argSet
				if _, ok := argSet[scanner.Text()]; ok {
					continue
				}

				// Skip if not a valid url
				if _, err := url.ParseRequestURI(scanner.Text()); err != nil {
					continue
				}

				for nonURL, curr := range nonURLSet {
					stripProtocol := regexp.MustCompile(`^(https?://)?(www\.)?`)
					similar := int8(utils.SimilarityScore(nonURL, stripProtocol.ReplaceAllString(scanner.Text(), "")))
					if similar == 0 {
						continue
					}

					// Add if lesser than 10
					if len(curr) < 10 {
						nonURLSet[nonURL][scanner.Text()] = similar
						continue
					}

					// Add if current has higher score
					lowestKey := ""
					for i, comp := range curr {
						if comp < similar {
							if lowestKey == "" || comp < curr[lowestKey] {
								lowestKey = i
							}
						}
					}

					if lowestKey != "" {
						nonURLSet[nonURL][lowestKey] = similar
					}
				}
			}

			Debug("Similar urls found in cache: " + fmt.Sprint(nonURLSet))

			// Prompt for proper url
			Info("Prompting user for proper URL...")
			for prev, comp := range nonURLSet {
				if len(comp) == 0 {
					fmt.Printf("No similar URLs found in cache for %s\n", prev)
					os.Exit(1)
				}

				ask := make([]string, 0, 10)
				mapping := map[int]string{}
				i := 0
				for key, score := range comp {
					ask = append(ask, fmt.Sprintf("%d. %s (%d%% match)", i, key, score))
					mapping[i] = key
					i++
				}

				var index int8
				response := utils.InputPrompt(utils.Multiline(append(ask, "", "Passed: "+prev, "Which url do you want to use? (Default: 0)")...))
				if response == "" {
					index = 0
				} else {
					i, err := strconv.ParseInt(response, 10, 8)
					if err != nil {
						fmt.Printf("%s is not a valid number\n", response)
						Debug(err.Error())
						os.Exit(1)
					}
					index = int8(i)
				}

				if index < 0 || index > int8(len(comp)) {
					fmt.Printf("%d is not a valid index\n", i)
					Debug(err.Error())
					os.Exit(1)
				}

				Info("Using url: " + mapping[int(index)])
				argSet[mapping[int(index)]] = struct{}{}
			}
		}

		// Ensure a URL was passed
		Debug("Final URL(s): " + fmt.Sprint(argSet))
		if len(argSet) == 0 {
			fmt.Println("No URL(s) were passed! See -h|--help for usage.")
			os.Exit(1)
		}

		// Caching URLS
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()
			fmt.Println("Writing URLs to cache file...")

			// Buffer writer
			buffer := bufio.NewWriter(file)
			for downloadURL := range argSet {
				if _, err := buffer.WriteString(downloadURL + "\n"); err != nil {
					fmt.Println("Failed to write to cache file at " + cacheFile)
					Debug(err.Error())
					return
				}
			}

			// Write to file
			if err := buffer.Flush(); err != nil {
				fmt.Println("Failed to write cache file at " + cacheFile)
				Debug(err.Error())
				return
			}

			fmt.Println("Wrote " + fmt.Sprint(len(argSet)) + " url(s) to cache file")
		}()

		downloadFile := func(url string) {
			// Get file
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

			split := strings.Split(url, "/")
			fileName := split[len(split)-1]

			if !strings.Contains(fileName, ".") {
				Info("Resolving file extension from content type...")
				contentType := response.Header.Get("Content-Type")
				extensions, err := mime.ExtensionsByType(contentType)
				if err != nil || len(extensions) == 0 {
					fmt.Println("Failed to resolve extension automatically for content type: " + contentType)
					if err != nil {
						Debug(err.Error())
					}
					return
				}

				fileName += extensions[0]
			}
			downloadLocation := filepath.Clean(filepath.Join(dirPath, fileName))

			// Ensure file already does not exist
			Info("Checking if " + downloadLocation + " already exists")
			if _, err := os.Stat(downloadLocation); err == nil {
				fmt.Printf("File already exists for %s\n", downloadLocation)
				Debug("File: " + downloadLocation + " already exists for " + url)
				return
			}

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
		}

		// Handle downloading
		switch getFlags.strategy {
		case utils.StrategyConcurrent:
			// Concurrency with no limit
			if getFlags.maxConcurrency == 0 {
				fmt.Printf("Downloading concurrently... [Use: %d, Max: No limit]\n", len(argSet))
				waitGroup.Add(len(argSet))

				for url := range argSet {
					go func(url string) {
						defer waitGroup.Done()
						downloadFile(url)
					}(url)
				}

				// Concurrency with limit
			} else {
				resolvedConcurrency := min(len(argSet), getFlags.maxConcurrency)
				fmt.Printf("Downloading concurrently... [Use: %d, Max: %d]\n", resolvedConcurrency, getFlags.maxConcurrency)

				// Establish channel and workers
				waitGroup.Add(resolvedConcurrency)
				ch := make(chan string)
				for t := 0; t < resolvedConcurrency; t++ {
					go func() {
						for url := range ch {
							downloadFile(url)
						}

						waitGroup.Done()
					}()
				}

				for url := range argSet {
					ch <- url
				}

				close(ch)
			}
		case utils.StrategySynchronous:
			fmt.Println("Downloading synchronously...")

			for url := range argSet {
				downloadFile(url)
			}
		}

		waitGroup.Wait()
	},
}

func init() {
	getFlags.strategy = utils.StrategyConcurrent

	rootCommand.AddCommand(getCommand)
	getCommand.Flags().IntVarP(&getFlags.maxConcurrency, "max-concurrency", "m", 10, "Maximum number of concurrent downloads [0 = unlimited] (default is 10)")

	getCommand.Flags().VarP(&getFlags.inputFile, "file", "f", "Path to the input file containing the url(s)")
	if err := getCommand.MarkFlagFilename("file"); err != nil {
		fmt.Println("Failed to mark flag -f as filename in get command")
		Debug(err.Error())
		os.Exit(1)
	}

	getCommand.Flags().VarP(&getFlags.strategy, "strategy", "s", "Strategy to use when downloading (default is concurrent)")
	if err := getCommand.RegisterFlagCompletionFunc("strategy", strategyCompletion); err != nil {
		fmt.Println("Failed to register completion for flag -s in get command")
		Debug(err.Error())
		os.Exit(1)
	}
}

func strategyCompletion(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"synchronous\tDownload videos sequentially",
		"concurrent\tDownload videos concurrently DEFAULT",
	}, cobra.ShellCompDirectiveNoFileComp
}
