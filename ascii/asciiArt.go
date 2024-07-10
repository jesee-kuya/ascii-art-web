package webart

import (
	"strings"
)

func Ascii(fileArr []string, words string) string {
	var result string

	words = strings.ReplaceAll(words, "\r\n", "\n")
	wordsArr := strings.Split(words, "\n")
	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					result += fileArr[int(start)+i]
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}
