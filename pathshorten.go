// Package pathshorten provides an implementation of Vim's "pathshorten"
// function, which shortens directory names in a file path.
package pathshorten

import (
	"os"
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
func PathShorten(input string, componentLength uint) string {
	separator := string(os.PathSeparator)
	tokens := strings.Split(input, separator)

	for i := 0; i < len(tokens); i++ {
		if i != len(tokens)-1 {
			tokens[i] = shortenPathComponent(tokens[i], componentLength)
		}
	}

	return strings.Join(tokens, separator)
}
