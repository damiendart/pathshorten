package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/damiendart/pathshorten"
)

func printOutput(output string, newline bool) {
	if newline {
		fmt.Println(output)
	} else {
		fmt.Print(output)
	}
}

func main() {
	pathComponentLength := flag.Uint(
		"length",
		1,
		"The number of alphanumeric characters of each directory to display",
	)
	suppressTrailingNewline := flag.Bool(
		"n",
		false,
		"Suppress trailing newline",
	)

	flag.Parse()

	if len(flag.Args()) > 0 {
		printOutput(
			pathshorten.PathShorten(flag.Args()[0], *pathComponentLength),
			!*suppressTrailingNewline,
		)
	} else {
		currentWorkingDirectory, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}

		printOutput(
			pathshorten.PathShorten(
				currentWorkingDirectory,
				*pathComponentLength,
			),
			!*suppressTrailingNewline,
		)
	}
}
