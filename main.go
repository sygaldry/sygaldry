package main

import (
	"fmt"
	"os"
	"strings"
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
		os.Exit(1)
	}
}

func suggestHelp(cliArg string) {
	fmt.Printf("%s: missing operand after `%s'\n", cliArg, cliArg)
	fmt.Printf("%s: Try `%s --help' for more information.\n", cliArg, cliArg)
	os.Exit(1)
}

func printHelp(cliArg string) {
	fmt.Println("Here's some help, bro!")
	os.Exit(1)
}
