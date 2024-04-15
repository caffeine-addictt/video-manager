package src

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/spf13/cobra"
)

var deleteFlags struct {
	yes bool
}

var deleteCommand = &cobra.Command{
	Use:     "delete <pattern>",
	Aliases: []string{"del", "rm"},
	Short:   "Delete videos from the source directory",
	Args:    cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		filenames, err := utils.ReadDirectory(workingDir, 0)
		if err != nil {
			fmt.Println("Failed to read source directory:", err.Error())
			os.Exit(1)
		}

		if toComplete == "" {
			return filenames, cobra.ShellCompDirectiveNoFileComp
		}

		for i := 0; i < len(filenames); i++ {
			match, err := regexp.MatchString(listFlags.query, filenames[i])
			if err != nil {
				return []string{"RegExp failed to match " + filenames[i] + ": " + err.Error()}, cobra.ShellCompDirectiveNoFileComp
			}

			if !match || strings.HasPrefix(filenames[i], toComplete) {
				filenames[len(filenames)-1], filenames[i], filenames = "", filenames[len(filenames)-1], filenames[:len(filenames)-1]
				i--
			}
		}

		return filenames, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		Info("Reading source directory")
		filenames, err := utils.ReadDirectory(workingDir, 0)
		if err != nil {
			fmt.Println("Failed to read source directory:", err.Error())
			os.Exit(1)
		}

		query := args[0]
		if query == "" {
			fmt.Println(utils.Multiline(
				"Argument '' passed is too vague.",
				"If you'd like to delete all files in the source directory, please pass the -a/--all flag or pass '.*' as the argument.",
				"",
				"See -h or --help for usage.",
			))
		}

		for i := 0; i < len(filenames); i++ {
			match, err := regexp.MatchString(query, filenames[i])

			if (err != nil && !match) || (err == nil && !strings.HasPrefix(filenames[i], query)) {
				Debug(fmt.Sprintf("%s does not contain or match %s", filenames[i], listFlags.query))
				filenames[len(filenames)-1], filenames[i], filenames = "", filenames[len(filenames)-1], filenames[:len(filenames)-1]
				i--
				continue
			}
		}

		if len(filenames) == 0 {
			fmt.Println("No matching files found.")
			os.Exit(1)
		}

		if !deleteFlags.yes {
			fmt.Println(strings.Join(filenames, "\n"))
			response := utils.InputPrompt("Delete these files? [y/N]")
			switch strings.ToLower(response) {
			case "y", "yes":
				break
			default:
				os.Exit(0)
			}
		}

		// Delete files
		for _, filename := range filenames {
			fileFullPath := filepath.Join(workingDir, filename)
			if err := os.Remove(fileFullPath); err != nil {
				fmt.Printf("Failed to delete %s: %s", filename, err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCommand.AddCommand(deleteCommand)
	deleteCommand.Flags().BoolVarP(&deleteFlags.yes, "yes", "y", false, "Skip confirmation prompt (defaults to false)")
}
