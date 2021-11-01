package output

import "ascii-art-web/models"

func Output(a *models.Art) (filename string) {
	if a.Page.FileName == "" {
		a.Page.FileName = "art_file"
	}

	if a.Page.FormatCurr != "txt" && a.Page.FormatCurr != "pdf" {
		a.Page.FormatCurr = "txt"
	}

	filename = a.Page.FileName + "." + a.Page.FormatCurr

	return filename
}
