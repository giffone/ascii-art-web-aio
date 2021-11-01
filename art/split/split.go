package split

import "ascii-art-web/models"

// SplitArgs checks for illegal symbols in arguments and split by new line
func SplitArgs(t *models.Text, text string) {
	t.Rune = nil
	for _, symbol := range text {
		if symbol == '\r' {
			continue
		} else {
			t.AddRune(symbol, 1)
		}
	}
}

// NewlineRune appending '\n' to rune slice
func NewLineRune(t *models.Text, make bool, count int) {
	if make {
		t.AddRune('\n', count)
	} else {
		lRune := len(t.Rune)
		if lRune > 0 {
			if t.Rune[lRune-1] != '\n' {
				t.AddRune('\n', count)
			}
		}
	}
}
