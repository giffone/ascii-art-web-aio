package alphabet

import (
	"ascii-art-web/art/check"
	"ascii-art-web/models"
	"fmt"
)

type Scan struct {
	Slice  []rune // scanned line (scan line by line) in rune
	LLine  int    // lenght of scanned line
	Symbol rune   // current symbol in rune
	Height int    // hight of each letter
	Code   int    // number of possition in ascii
	Lines  int    // lines in txt file
}

// MakeAlphabet makes map of ascii-art
func MakeAlphabet(a *models.Art, b *models.Buf) error {
	if _, ok := a.Alphabet.Rune[b.Title]; !ok {
		a.Alphabet.Rune[b.Title] = make(map[rune][]string)
	}

	// read scanned bufer
	if err := scanRune(a, b); err != nil {
		return err
	}
	// if alphabet does not consist 95 symbols (32-126)
	lenght := len(a.Alphabet.Rune[b.Title])
	if lenght != 95 {
		return fmt.Errorf("alphabet: %s there is no 95 "+
			"letters imported, just: %d", b.Title, lenght)
	}
	// checks for missing runes
	if err := check.Maps(a, b); err != nil {
		return err
	}
	return nil
}

// scanRune make map with key - rune
func scanRune(a *models.Art, b *models.Buf) error {
	// first symbol -32 and width of each symbol - exmpl 9
	s := &Scan{Code: 32, Height: b.Height}
	for _, line := range b.ArrStr1 {
		r := []rune(line)
		s.LLine = len(r)
		if s.LLine != 0 {
			s.Symbol = rune(s.Code)
			// save rune exept last line ""
			a.Alphabet.Rune[b.Title][s.Symbol] = append(a.Alphabet.Rune[b.Title][s.Symbol], line)
			if s.Height != s.Lines {
				// save lenght of each symbol, if not exist, checks for errors
				if err := lenScanRune(a, s, b); err != nil {
					return err
				}
				// go to next map
			} else {
				s.Code++
				s.Height += b.Height + 1 // correct +1 a newline between symbols
			}
		}
		s.Lines++
	}
	return nil
}

// lenScanRune saves lenght of symbol, if not exist, and compare for each line, if compares is different - error
func lenScanRune(a *models.Art, s *Scan, b *models.Buf) error {
	if _, ok := a.Alphabet.RuneLen[s.Symbol]; !ok {
		a.Alphabet.RuneLen[s.Symbol] = s.LLine
	} else {
		if s.LLine != a.Alphabet.RuneLen[s.Symbol] {
			return fmt.Errorf("alphabet: %s length of symbol "+
				"\"%q\" is different.", b.Title, s.Symbol)
		}
	}
	return nil
}
