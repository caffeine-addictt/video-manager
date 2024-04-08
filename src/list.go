package src

import (
	"fmt"
	"os"
	"strings"

	"github.com/caffeine-addictt/video-manager/src/utils"
	"github.com/spf13/cobra"
)

var listFlags struct {
	allowExt  []string
	rejectExt []string
	count     int
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List the downloaded videos",
	Long:  "List the downloaded videos in the specified directory",
	Run: func(cmd *cobra.Command, args []string) {
		considerAllow := len(listFlags.allowExt) > 0
		considerReject := len(listFlags.rejectExt) > 0

		Info("Reading directory...")
		filenames, err := utils.ReadDirectory(workingDir, listFlags.count)
		if err != nil {
			fmt.Println("Failed to read directory")
			Debug(err.Error())
			os.Exit(1)
		}

		// Merge accept and reject extensions
		extMap := make(map[string]struct{})
		for _, allow := range listFlags.allowExt {
			extMap[allow] = struct{}{}
		}
		for _, reject := range listFlags.rejectExt {
			extMap[reject] = struct{}{}
		}
		Debug(fmt.Sprintf("Merged extensions: %v", extMap))

		Info("Filtering extensions...")
		if considerAllow || considerReject {
			for i := 0; i < len(filenames); i++ {
				split := strings.Split(filenames[i], ".")
				extension := split[len(split)-1]

				_, ok := extMap[extension]
				if (considerAllow && !ok) || (considerReject && ok) {
					filenames[len(filenames)-1], filenames[i], filenames = "", filenames[len(filenames)-1], filenames[:len(filenames)-1]
					i--
				}
			}
		}

		fmt.Println(utils.Multiline(filenames...))
	},
}

func init() {
	rootCommand.AddCommand(listCommand)
	listCommand.Flags().IntVarP(&listFlags.count, "count", "n", 0, "Number of videos to list [0 = all] (defaults to 0)")
	listCommand.Flags().StringSliceVarP(&listFlags.allowExt, "allow", "a", []string{}, "The extensions of the videos to list [mp4 etc.] (defaults to [])")
	listCommand.Flags().StringSliceVarP(&listFlags.rejectExt, "exclude", "e", []string{}, "The extensions of the videos to list exclude [mp4 etc] (defaults to [])")

	listCommand.MarkFlagsMutuallyExclusive("allow", "exclude")
}
