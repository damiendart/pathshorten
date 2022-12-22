// Copyright (C) 2022 Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

// Package pathshorten provides an implementation of Vim's "pathshorten"
// function, which shortens directory names in a file path.
package pathshorten

import (
	"strings"
)

func shortenPathComponent(input string, length uint) string {
	runes := []rune(input)

	for i, item := range runes {
		if item == '.' || item == '~' {
			continue
		}

		if i+int(length) > len(runes) {
			return string(runes)
		}

		return string(runes[0 : i+int(length)])
	}

	return input
}

// PathShorten shortens directory names in a file path.
func PathShorten(
	input string,
	pathSeparator string,
	pathComponentLength uint,
) string {
	tokens := strings.Split(input, pathSeparator)

	for i := 0; i < len(tokens); i++ {
		if i != len(tokens)-1 {
			tokens[i] = shortenPathComponent(tokens[i], pathComponentLength)
		}
	}

	return strings.Join(tokens, pathSeparator)
}
