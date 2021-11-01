// package structure strores all data structures
package models

import (
	"net/http"
)

type Art struct {
	Alphabet Alphabet
	Color    Color
	Text     Text
	Page     Page
}

// MenuFont appends list of fonts for web menu
func (p *Art) MenuFont(s string) {
	p.Page.Fonts = append(p.Page.Fonts, s)
}

// SavePrevious appends maked Art to web <pre>
func (a *Art) SavePrevious(b *Buf) {
	if a.Page.AnsSave == "on" {
		b.ArrStr1 = append(append(b.ArrStr1, " "), a.Page.AnsWeb...)
		b.ArrStr2 = append(append(b.ArrStr2, " "), a.Page.AnsFile...)
	}
	a.Page.AnsWeb = b.ArrStr1
	a.Page.AnsFile = b.ArrStr2
}

type Page struct {
	TextArea string   // user input field
	Status   []string // text for error
	Menu
	Save
	Ans
}

func (p *Page) AddColors() {
	p.Colors = []string{"", "black", "red", "green", "yellow",
		"blue", "purple", "cyan", "white", "orange"}
}

func (p *Page) AddFormats() {
	p.Formats = []string{"txt", "pdf"}
}

func (a *Page) WriteStatus(s string) {
	a.Status = append(a.Status, s)
}

func (p *Page) Add(r *http.Request, form string) {
	if form == "submit" {
		p.TextArea = r.PostFormValue("textarea")
		p.FontCurr = r.PostFormValue("fonts")
		p.ColorCurr = r.PostFormValue("colors")
		p.ColorParam = r.PostFormValue("colorparam")
		p.AnsSave = r.PostFormValue("anssave")
	} else if form == "save" {
		p.FileName = r.PostFormValue("filename")
		p.FormatCurr = r.PostFormValue("formats")
	}
}

func (p *Page) Clear() {
	p.TextArea = ""
	p.FontCurr = ""
	p.ColorCurr = ""
	p.ColorParam = ""
	p.AnsSave = ""
	p.FileName = ""
	p.FormatCurr = ""
	p.AnsWeb = nil
	p.AnsFile = nil
	p.Status = nil
}

type Menu struct {
	Fonts      []string // menu fonts
	FontCurr   string   // choosen font
	Colors     []string // menu color
	ColorCurr  string   // choosen color
	ColorParam string   // menu with parameters coloring
}

type Save struct {
	FileName   string   // filename
	Formats    []string // format of file (txt or pdf)
	FormatCurr string   // choosen format
}

type Ans struct {
	AnsWeb  []string // maked Art for web <pre>
	AnsFile []string   // maked Art for save into file
	AnsSave string   // checkbox if need save art in <pre> and add into new
}

type Alphabet struct {
	Sourse   []string                     // scanned file
	Rune     map[string]map[rune][]string // sorted in ascii code
	RuneLen  map[rune]int                 // save lenght of each symbol
	Letter   rune                         // current letter
	Coloring bool                         // if need to color
}

type Text struct {
	Rune []rune // user splited text from textarea
}

// AddRune appends rune in []rune if needs
func (t *Text) AddRune(r rune, count int) {
	for count != 0 {
		t.Rune = append(t.Rune, r)
		count--
	}
}

type Color struct {
	MethodColoring string // method coloring: and (a-d), or (a,b,c), none, all
	MethodBy       string // by indexes or letters
	Case1          string // color start
	Case2          string // color end (reset)
	BySymbol       BySymbol
	ByIndex        ByIndex
}

type BySymbol struct {
	Range  []string // to save letters (a-d), (a,b,c)
	Range1 rune     // first letter
	Range2 rune     // last letter
}

type ByIndex struct {
	Range    []int // to save indexes [1:4], [0,3,8]
	Range1   int   // first index
	Range2   int   // last index
	MaxIndex bool  // lenght of all letters in arguments (to make max index)
}

type Buf struct {
	ArrStr1 []string
	ArrStr2 []string
	Title   string
	Height  int
	Lines   int
	Index   int
}

func (b *Buf) AddArrStr1(text string) {
	b.ArrStr1 = append(b.ArrStr1, text)
}
