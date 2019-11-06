/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sygaldry/sygaldry-core/stage"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sygaldry",
	Short: "Usage: sygaldry <step name> -f <rune yaml file>",
	Long:  `Usage: sygaldry <step name> -f <rune yaml file>`,
	Run: func(cmd *cobra.Command, args []string) {
		file := cmd.Flag("file")

		if !strings.HasSuffix(file.Value.String(), ".yaml") && !strings.HasSuffix(file.Value.String(), ".yml") {
			fmt.Printf("Your runes file should be a yaml.")
		} else if len(args) == 1 {
			stage.GetStage(file.Value.String(), args[0])
		} else {
			fmt.Println("Usage: sygaldry <step name> -f <rune yaml file>")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sygaldry.yaml)")

	rootCmd.Flags().StringP("file", "f", "", "rune yaml file to parse")
	rootCmd.MarkFlagRequired("file")
}
