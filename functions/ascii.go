package web

import "strings"

func Ascii(input string, fileArr []string) string {
	input = Reap(input)
	arr := strings.Split(input, "\r")
	var count int
	var outputBuilder strings.Builder

	for _, val := range arr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					outputBuilder.WriteString(fileArr[int(start)+i])
				}
				outputBuilder.WriteString("\r")
			}
		} else {
			count++
			if count < len(arr) {
				outputBuilder.WriteString("\r")
			}
		}
	}
	return outputBuilder.String()
}
