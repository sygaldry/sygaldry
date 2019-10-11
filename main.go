package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/smallfish/simpleyaml"
)

/*
	I am sure there's a better way to do this, I'm just trying to get the hang of Go right now.
	Edvin probably has something better for CLI stuff already.
*/
func main() {

	// We always get this arg.
	cliArg := strings.ReplaceAll(os.Args[0], "./", "")

	// Base case we get exactly one arg, we want to point the user to the help
	if len(os.Args) == 1 {
		suggestHelp(cliArg)
	}

	if len(os.Args) == 2 || len(os.Args) == 3 {
		firstArg := os.Args[1]
		if firstArg == "--help" {
			printHelp(cliArg)
		} else {
			suggestHelp(cliArg)
		}
	}

	if len(os.Args) >= 4 {
		runeArg := os.Args[1]
		forceFlagArg := os.Args[2]
		fileArg := os.Args[3]

		if runeArg == "--help" {
			printHelp(cliArg)
		}

		if forceFlagArg != "-f" {
			fmt.Printf("Did you mean this?\nsygaldry %s -f %s\n", runeArg, forceFlagArg)
			os.Exit(1)
		}

		fmt.Printf("Attempting to run the %s step in %s\n", runeArg, fileArg)
		parseYaml(fileArg, runeArg)
		os.Exit(1)
	}
}

// This can certainly be done better and be split up a bunch more but it's a start!
func parseYaml(fileArg string, runeArg string) {
	source, err := ioutil.ReadFile(fileArg)
	if err != nil {
		panic(err)
	}

	yaml, err := simpleyaml.NewYaml(source)
	if err != nil {
		panic(err)
	}

	yam, err := yaml.Get(runeArg).Map()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", yam)

	runeMap := yaml.Get(runeArg)
	runeKeys, err := runeMap.GetMapKeys()
	if err != nil {
		panic(err)
	}
	mainRune := runeKeys[0]
	fmt.Printf("\nTargeting rune: %s\n", mainRune)

	argMap := runeMap.Get(mainRune)
	argKeys, err := argMap.GetMapKeys()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nPushing params:")
	for _, key := range argKeys {
		val, err := argMap.Get(key).String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("    %s:\t%s\n", key, val)
	}
	fmt.Println("")
}

// We can suggest help but there's not much there yet.
func suggestHelp(cliArg string) {
	fmt.Printf("%s: missing operand after `%s'\n", cliArg, cliArg)
	fmt.Printf("%s: Try `%s --help' for more information.\n", cliArg, cliArg)
	os.Exit(1)
}

// There's not much help to really be had here...
func printHelp(cliArg string) {
	fmt.Println("Here's some help, bro! <3")
	os.Exit(1)
}
