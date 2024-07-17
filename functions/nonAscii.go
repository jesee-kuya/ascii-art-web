package web

import "fmt"

func NonAscii(input string) error {
	for _, v := range input {
		if v > 126 {
			return fmt.Errorf(" Bad request")
		}
	}
	return nil
}
