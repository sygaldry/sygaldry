package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sygaldry/sygaldry-core/stage"
)

// stageCmd represents the base command when called without any subcommands
var stageCmd = &cobra.Command{
	Use:     "sygaldry",
	Short:   "Usage: sygaldry <stage name> -f <[path | URL] rune yaml file>",
	Long:    `Usage: sygaldry <stage name> -f <[path | URL] rune yaml file>`,
	Example: "sygaldry build -f runes.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		file := cmd.Flag("file")

		if !strings.HasSuffix(file.Value.String(), ".yaml") && !strings.HasSuffix(file.Value.String(), ".yml") {
			fmt.Printf("Runes file needs to be a yaml file\n")
		} else if len(args) == 1 {
			currentStage, getStageErr := stage.GetStage(file.Value.String(), args[0])
			if getStageErr != nil {
				panic(getStageErr)
			}
			runStageErr := currentStage.Run()
			if runStageErr != nil {
				panic(runStageErr)
			}
		} else {
			fmt.Println("Usage: sygaldry <stage name> -f <[path | URL] rune yaml file>")
		}
	},
}

// Execute Method for executing stage command
func Execute() {
	if err := stageCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	stageCmd.Flags().StringP("file", "f", "", "rune yaml file to parse")
	stageCmd.MarkFlagRequired("file")
}
