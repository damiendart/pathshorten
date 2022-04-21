package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/damiendart/pathshorten"
)

func main() {
	var path string

	pathComponentLength := flag.Uint(
		"length",
		1,
		"The number of alphanumeric characters of each directory to display",
	)
	pathSeparator := flag.String(
		"separator",
		string(os.PathSeparator),
		"The path separator",
	)
	suppressTrailingNewline := flag.Bool(
		"n",
		false,
		"Suppress trailing newline",
	)

	flag.Parse()

	if len(flag.Args()) > 0 {
		path = flag.Args()[0]
	} else {
		currentWorkingDirectory, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}

		path = currentWorkingDirectory
	}

	output := pathshorten.PathShorten(
		path,
		*pathSeparator,
		*pathComponentLength,
	)

	if *suppressTrailingNewline {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}
