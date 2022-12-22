/*
Pathshorten shortens directory names in a file path.
It is inspired by Vim's "pathshorten" function.

Usage:

	pathshorten [flags] [path]

Arguments:

	path string
		The file path to shorten. The path does not have to exist. If
		not provided, the current working directory is used.

Optional flags:

	-length unit
		The number of alphanumeric characters of each directory to
		display, The default is one character.
	-n
		Suppress adding a newline to the end of any non-error output.
	-separator string
		The path separator to use when splitting the file path into
		individual components. The default is the path separator used by
		the operating system.
*/
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
		"the number of alphanumeric characters of each directory to display",
	)
	pathSeparator := flag.String(
		"separator",
		string(os.PathSeparator),
		"the path separator to use when splitting the file path into individual components",
	)
	suppressTrailingNewline := flag.Bool(
		"n",
		false,
		"suppress adding a newline to the end of any non-error output",
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
