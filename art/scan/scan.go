package scan

import (
	"ascii-art-web/art/alphabet"
	"ascii-art-web/art/fonts"
	"ascii-art-web/models"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Scan struct {
	file *os.File
}

func (s *Scan) open(name string) error {
	var err error

	s.file, err = os.Open(name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scan) scan(a *models.Art, b *models.Buf) error {
	defer s.file.Close()

	sourse := bufio.NewScanner(s.file)

	if err := sourse.Err(); err != nil {
		return err
	}

	lines := 0
	for sourse.Scan() {
		b.AddArrStr1(sourse.Text())
		lines++
	}

	// if file does not consist 855 lines in fontfile
	if lines != b.Lines {
		return fmt.Errorf("scan: something wrong with font file: %d"+
			"lines instead of %d", lines, b.Lines)
	}
	return nil
}

func Directory(a *models.Art, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("Empy dir.")
		return
	}

	a.Alphabet.Rune = make(map[string]map[rune][]string)

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".txt") {
			s := new(Scan)
			b := new(models.Buf)
			if err := s.open(dir + "/" + name); err == nil {
				b.Title = strings.TrimSuffix(name, ".txt")
				fonts.FontParam(b)
				if err := s.scan(a, b); err == nil {
					a.Alphabet.RuneLen = make(map[rune]int)
					if err := alphabet.MakeAlphabet(a, b); err == nil {
						a.MenuFont(b.Title)
					} else {
						log.Println(err.Error())
						a.Alphabet.Rune[b.Title] = nil
					}
				} else {
					log.Println(err.Error())
				}
			} else {
				log.Println(err.Error())
			}
		}
		a.Alphabet.RuneLen = nil
	}
}
