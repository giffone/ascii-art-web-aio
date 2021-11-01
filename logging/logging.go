package logging

import (
	"ascii-art-web/models"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PrintOut(s []string) {
	for _, line := range s {
		log.Println(line)
	}
}

func WarningBrc(a *models.Art, value string) {
	bracks := 2
	for _, letter := range value {
		if letter == '[' || letter == ']' {
			bracks--
		}
	}
	if bracks < 2 {
		a.Page.WriteStatus(fmt.Sprintf("color: \"%s\" parameter is wrong.", value))
	}
}

func WarningInd(a *models.Art, count int) {
	a.Page.WriteStatus(fmt.Sprintf("color: %d wrong indexes skipped.", count))
}

func WarningErr(a *models.Art, err error, s string) {
	a.Page.WriteStatus(s)
	log.Println(err)
}

// formatRequest generates ascii representation of a request
func FormatRequest(r *http.Request) string {
	fmt.Println()
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	// for name, headers := range r.Header {
	// 	name = strings.ToLower(name)
	// 	for _, h := range headers {
	// 		request = append(request, fmt.Sprintf("%v: %v", name, h))
	// 	}
	// }

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
