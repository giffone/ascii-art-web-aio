package input

import (
	chk "ascii-art-web/art/check"
	cl "ascii-art-web/art/colored"
	"ascii-art-web/models"
	"html"
	"strings"
)

// MakeArt beagin to make art
func MakeArt(a *models.Art, b *models.Buf) {
	// choose color start / end (reset)
	cl.ChooseColor(a)
	if a.Color.MethodBy == "bySymbol" {
		if a.Color.MethodColoring == "and" {
			// coloring in range between letter1 and letter2
			a.Color.BySymbol.Range1, a.Color.BySymbol.Range2 = findAndSymbol(a)
		}
	} else if a.Color.MethodBy == "byIndex" {
		// add max number for range if parameter was [5:]
		if a.Color.ByIndex.MaxIndex {
			lRune := len(a.Text.Rune)
			a.Color.ByIndex.Range = append(a.Color.ByIndex.Range, lRune)
			a.Color.MethodColoring = "and"
		}
		if a.Color.MethodColoring == "and" {
			// coloring in range between letter1 and letter2
			a.Color.ByIndex.Range1, a.Color.ByIndex.Range2 = findAndIndex(a)
		}
	}
	// read letters from argument
	letters(a, b)
}

// letters finds letters of the word in the standard.txt banner and prints it
func letters(a *models.Art, b *models.Buf) {
	// count text blocks
	l := chk.Lines(a, b)
	b.ArrStr1 = make([]string, l)
	b.ArrStr2 = make([]string, l)
	b.Index = b.Height
	// if method coloring all words, add color to the begining of the line
	if a.Color.MethodColoring == "all" {
		cl.AddStartColor(a, b)
	}

	index := 0
	for _, symbol := range a.Text.Rune {
		if symbol == '\r' {
			continue
		}
		if symbol == '\n' {
			if a.Alphabet.Letter != 0 && a.Alphabet.Letter != '\n' {
				b.Index += b.Height + 1 // correct +1 newline
			} else {
				b.Index += 1
			}
			a.Alphabet.Letter = symbol
			continue
		}
		// remember current letter
		a.Alphabet.Letter = symbol
		// by default coloring is false
		a.Alphabet.Coloring = false
		// coloring for letter by letter
		if symbol != ' ' && symbol != '\t' {
			if a.Color.MethodColoring == "or" {
				if a.Color.MethodBy == "bySymbol" {
					if findOrSymbol(a, symbol) {
						// coloring in range between one letter
						a.Color.BySymbol.Range1, a.Color.BySymbol.Range2 = symbol, symbol
					} else {
						// clear range for the next one
						a.Color.BySymbol.Range1, a.Color.BySymbol.Range2 = 0, 0
					}
				} else if a.Color.MethodBy == "byIndex" {
					if findOrIndex(a, index) {
						// coloring in range between one letter
						a.Color.ByIndex.Range1, a.Color.ByIndex.Range2 = index, index
					} else {
						// clear range for the next one
						a.Color.ByIndex.Range1, a.Color.ByIndex.Range2 = -1, -1
					}
				}
			}
			if a.Color.MethodBy == "bySymbol" {
				// if choosen letter in range - make color
				if symbol >= a.Color.BySymbol.Range1 && symbol <= a.Color.BySymbol.Range2 {
					a.Alphabet.Coloring = true
				}
			} else if a.Color.MethodBy == "byIndex" {
				// if choosen letter in range - make color
				if index >= a.Color.ByIndex.Range1 && index <= a.Color.ByIndex.Range2 {
					a.Alphabet.Coloring = true
				}
			}
			index++
		}
		makeOutputByMethod(a, b)
	}
}

// makeOutputByMethod make output by method
func makeOutputByMethod(a *models.Art, b *models.Buf) {
	if a.Color.MethodColoring == "none" || a.Color.MethodColoring == "all" {
		// if method "all" - it is no need to add color for each letter. only at start and the end line
		makeOutput(a, b, false)
	} else {
		makeOutput(a, b, a.Alphabet.Coloring)
	}
}

// makeOutput append to array founded letter
func makeOutput(a *models.Art, b *models.Buf, coloring bool) {
	index := 0
	for i := b.Index - b.Height; i < b.Index; i++ {
		var tmp strings.Builder
		line := a.Alphabet.Rune[b.Title][a.Alphabet.Letter][index]
		line2 := html.EscapeString(line)
		if coloring {
			tmp.WriteString(a.Color.Case1)
			tmp.WriteString(line2)
			tmp.WriteString(a.Color.Case2)
			b.ArrStr1[i] += tmp.String()
			// b.ArrStr1[i] += a.Color.Case1 + line + a.Color.Case2
		} else {
			b.ArrStr1[i] += line2
		}
		b.ArrStr2[i] += line
		index++
	}
}

// findAndSymbol find "a" and "f" letters in range (a-f)
func findAndSymbol(a *models.Art) (symbol1 rune, symbol2 rune) {
	for _, word := range a.Color.BySymbol.Range {
		for _, symbol := range word {
			if symbol1 == 0 {
				symbol1 = symbol
			}
			symbol2 = symbol
		}
	}
	// if wrong range, fix it
	if symbol1 > symbol2 {
		symbol1, symbol2 = symbol2, symbol1
	}
	return symbol1, symbol2
}

// findAndIndex find "0" and "8" indexes in range [0-8]
func findAndIndex(a *models.Art) (index1 int, index2 int) {
	sourse := a.Color.ByIndex.Range
	lSourse := len(sourse)

	index1 = sourse[0]
	index2 = sourse[lSourse-1]
	// if wrong range, fix it
	if index1 > index2 {
		index1, index2 = index2, index1
	}
	return index1, index2
}

// findOrSymbol find "a" or "f" letters in range (a,f)
func findOrSymbol(a *models.Art, symbol rune) bool {
	for _, word := range a.Color.BySymbol.Range {
		for _, letterF := range word {
			if symbol == letterF {
				return true
			}
		}
	}
	return false
}

// findOrIndex find "0" or "8" indexes in range (0,8)
func findOrIndex(a *models.Art, index int) bool {
	for _, indexF := range a.Color.ByIndex.Range {
		if index == indexF {
			return true
		}
	}
	return false
}
