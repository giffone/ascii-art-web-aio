package fonts

import "ascii-art-web/models"

func FontParam(b *models.Buf) {
	font := b.Title
	if font == "standard" {
		b.Height = 8
		b.Lines = 855
	} else if font == "thinkertoy" {
		b.Height = 8
		b.Lines = 855
	} else if font == "shadow" {
		b.Height = 8
		b.Lines = 855
	} else if font == "acrobatic" {
		b.Height = 12
		b.Lines = 1235
	} else if font == "lockergnome" {
		b.Height = 4
		b.Lines = 475
	} else if font == "marquee" {
		b.Height = 8
		b.Lines = 855
	} else if font == "wavy" {
		b.Height = 4
		b.Lines = 475
	} else if font == "straight" {
		b.Height = 4
		b.Lines = 475
	} else if font == "zigzag" {
		b.Height = 8
		b.Lines = 855
	}
}
