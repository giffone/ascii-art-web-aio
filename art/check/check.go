// package check contains all functions related to data validation
package check

import (
	"ascii-art-web/models"
	"fmt"
)

// AsciiSymbols checks for illegal symbols in arguments
func AsciiSymbols(text string) bool {
	for _, symbol := range text {
		if symbol == '\n' || symbol == '\r' {
			continue
		}
		// searches symbol between " "(space) and "~"(tilda) in ASCII table, in another case exits
		if symbol < ' ' || symbol > '~' {
			return false
		}
	}
	return true
}

// Brackets checks if value has two bracket opened and closed [...] or {...}
func Brackets(value string, lValue int, brackets ...string) bool {
	count := 0
	if lValue > 1 {
		for _, bracket := range brackets {
			if bracket[0] == value[0] && bracket[1] == value[lValue-1] {
				count++
			}
		}
	}
	return count != 0
}

// ByteConsist checks if the value is consist in []byte
func ByteConsist(value byte, symbols ...byte) bool {
	count := 0
	for _, symbol := range symbols {
		if symbol == value {
			count++
		}
	}
	return count != 0
}

// Maps checks for missing ascii symbols in created maps
func Maps(abc *models.Art, b *models.Buf) error {
	alphabet := []rune{' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
		+'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		+':', ';', '<', '=', '>', '?', '@',
		+'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		+'[', '\\', ']', '^', '_', '`',
		+'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		+'{', '|', '}', '~'}

	for _, symbol := range alphabet {
		_, ok := abc.Alphabet.Rune[b.Title][symbol]
		if !ok {
			return fmt.Errorf("alphabet: symbol \"%q\""+
				"is missing in %s", symbol, b.Title)
		}
	}
	return nil
}

// Lines checks if line (text) empty or not
func Lines(abc *models.Art, b *models.Buf) (count int) {
	text := false
	for _, symbol := range abc.Text.Rune {
		if symbol == '\n' {
			if text {
				count += b.Height + 1 // with newline
				text = false
				continue
			}
			count++
		} else {
			text = true
		}
	}
	return count
}
