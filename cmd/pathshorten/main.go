// Copyright (C) Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

/*
Pathshorten shortens directory names in a file path.
It is inspired by Vim's "pathshorten" function.

Usage:

	pathshorten [flags] [paths]

Arguments:

	paths
		File paths to shorten. The path does not have to exist. If
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
	"strings"

	"github.com/damiendart/pathshorten"
)

func main() {
	var paths []string

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
		paths = append(paths, flag.Args()...)
	} else {
		currentWorkingDirectory, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}

		paths = append(paths, currentWorkingDirectory)
	}

	for i, path := range paths {
		paths[i] = pathshorten.PathShorten(
			path,
			*pathSeparator,
			*pathComponentLength,
		)
	}

	if *suppressTrailingNewline {
		fmt.Print(strings.Join(paths, "\n"))
	} else {
		fmt.Println(strings.Join(paths, "\n"))
	}
}
