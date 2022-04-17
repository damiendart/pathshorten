package pathshorten

import (
	"os"
	"strings"
)

func shortenToken(input string, length uint) string {
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

func PathShorten(input string, tokenLength uint) string {
	separator := string(os.PathSeparator)
	tokens := strings.Split(input, separator)

	for i := 0; i < len(tokens); i++ {
		if i != len(tokens)-1 {
			tokens[i] = shortenToken(tokens[i], tokenLength)
		}
	}

	return strings.Join(tokens, separator)
}
