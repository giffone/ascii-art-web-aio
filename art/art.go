// Ascii-art all in one (default, fs, color, output, justify, reverse)
package art

import (
	"ascii-art-web/art/check"
	"ascii-art-web/art/colored"
	"ascii-art-web/art/fonts"
	"ascii-art-web/art/input"
	"ascii-art-web/art/split"
	"ascii-art-web/models"
)

func ArtMain(a *models.Art, b *models.Buf) {
	text := a.Page.TextArea
	b.Title = a.Page.FontCurr
	fonts.FontParam(b)

	a.Page.Status = nil

	if !check.AsciiSymbols(text) {
		a.Page.WriteStatus("ascii: not correct symbols in text")
		return
	}
	// find parameters of coloring if exist
	colored.Find(a)
	// check symbols, "\n" and append into slice
	if text == "" {
		split.NewLineRune(&a.Text, true, 1)
	} else {
		split.SplitArgs(&a.Text, text)
		split.NewLineRune(&a.Text, false, 1)
	}
	// log.Println(a.Text.Rune)
	input.MakeArt(a, b)
	a.SavePrevious(b)

	// logging.PrintOut(a.Page.AnsWeb)
	// logging.PrintOut(a.Page.AnsFile)
}
