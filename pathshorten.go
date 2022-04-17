package pathshorten

import (
	"os"
	"strings"
)

func shortenToken(input string) string {
	runes := []rune(input)

	for i, item := range runes {
		if item == '.' || item == '~' {
			continue
		}

		return string(runes[0 : i+1])
	}

	return input
}

func PathShorten(input string) string {
	separator := string(os.PathSeparator)
	tokens := strings.Split(input, separator)

	for i := 0; i < len(tokens); i++ {
		if i != len(tokens)-1 {
			tokens[i] = shortenToken(tokens[i])
		}
	}

	return strings.Join(tokens, separator)
}
