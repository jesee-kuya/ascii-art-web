package web

import (
	"fmt"
	"strings"
)

func NonAscii(input string) error {
	input = strings.ReplaceAll(input, "\r\n", "")
	for _, v := range input {
		if v > 126 || v < 32 {
			return fmt.Errorf(" Bad request")
		}
	}
	return nil
}
