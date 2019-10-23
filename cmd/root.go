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
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/smallfish/simpleyaml"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sygaldry",
	Short: "Usage: sygaldry <step name> -f <rune yaml file>",
	Long:  `Usage: sygaldry <step name> -f <rune yaml file>`,
	Run: func(cmd *cobra.Command, args []string) {
		file := cmd.Flag("file")

		if !strings.HasSuffix(file.Value.String(), ".yaml") || !strings.HasSuffix(file.Value.String(), ".yml") {
			fmt.Printf("Your runes file should be a yaml.")
		} else if len(args) == 1 && strings.HasSuffix(file.Value.String(), ".yaml") {
			parseYaml(args[0], file.Value.String())
		} else {
			fmt.Println("Usage: sygaldry <step name> -f <rune yaml file>")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sygaldry.yaml)")

	rootCmd.Flags().StringP("file", "f", "", "rune yaml file to parse")
	rootCmd.MarkFlagRequired("file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".sygaldry" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sygaldry")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func parseYaml(stage string, runeYaml string) {
	source, err := ioutil.ReadFile(runeYaml)
	if err != nil {
		log.Fatalln(err)
	}

	yaml, err := simpleyaml.NewYaml(source)
	if err != nil {
		log.Fatalln("2", err)
	}

	rune := yaml.Get(stage)
	runeKeys, err := rune.GetMapKeys()
	if err != nil {
		log.Fatalln("Step does not exist in provided rune yaml.")
	}
	mainRune := runeKeys[0]
	fmt.Printf("\nTargeting rune: %s\n", mainRune)

	argMap := rune.Get(mainRune)
	argKeys, err := argMap.GetMapKeys()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nPushing params:")
	for _, key := range argKeys {
		val, err := argMap.Get(key).String()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("    %s:\t%s\n", key, val)
	}
	fmt.Println("")
}
