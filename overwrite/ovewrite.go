package overwrite

import (
	"strings"
	"unicode/utf8"
)

func OverwriteString(str string, value rune, n int) string {
	// caso 1: deveria ser n >= len(str)
	// caso 2: deveria ser n >= utf8.RuneCountInString(str)
	// len string retorna o número de bytes, não de bytes em uma string
	if n >= utf8.RuneCountInString(str) {
		return strings.Repeat(string(value), len(str))
	}

	result := []rune(str)
	for i := 0; i <= n; i++ {
		result[i] = value
	}
	return string(result)
}
