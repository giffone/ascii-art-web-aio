package router

import (
	"ascii-art-web/art"
	"ascii-art-web/logging"
	"ascii-art-web/models"
	"ascii-art-web/output"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

type HandlerArt struct {
	Art *models.Art
}

var (
	templates *template.Template
	validPath *regexp.Regexp
)

func init() {
	templates = template.Must(template.ParseFiles("assets/template/index.html"))
	validPath = regexp.MustCompile("^/(ascii-art|download|error|favicon.ico)$")
}

func NewHandlerArt() *HandlerArt {
	memory := new(models.Art)
	return &HandlerArt{memory}
}

func (h *HandlerArt) mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed) // 405
		log.Println(logging.FormatRequest(r))
		return
	}
	if r.URL.Path != "/" {
		log.Printf("pattern is %s", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.URL.Path == "favicon.ico" {
		w.Header().Set("Content-Type", "image/x-icon")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		http.ServeFile(w, r, "assets/ico/favicon.ico")
	}
	log.Println(logging.FormatRequest(r))
	renderTemplate(w, "index", &h.Art.Page)
}

func (h *HandlerArt) artHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b := new(models.Buf)

		h.Art.Page.Add(r, "submit")
		art.ArtMain(h.Art, b)
		log.Println(logging.FormatRequest(r))
		w.WriteHeader(http.StatusCreated)
	}
	renderTemplate(w, "index", &h.Art.Page)
	// if r.Method == "GET" {
	// 	renderTemplate(w, "index", &h.Art.Page)
	// }
}

func (h *HandlerArt) downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.Art.Page.Add(r, "save")

		filename := output.Output(h.Art)
		arr := []byte{}

		for _, line := range h.Art.Page.AnsFile {
			arr = append(arr, []byte(line)...)
			arr = append(arr, '\n')
		}
		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(arr)))

		if _, err := w.Write(arr); err != nil {
			logging.WarningErr(h.Art, err, err.Error())
			return
		}
		log.Println(logging.FormatRequest(r))
	}
	// if r.Method == "GET" {
	// 	renderTemplate(w, "index", &h.Art.Page)
	// }
}

func checkHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("pattern is %s", r.URL.Path)
		if r.URL.Path == "/" {
			fn(w, r)
		} else {
			ptrn := validPath.FindStringSubmatch(r.URL.Path)
			if ptrn == nil {
				log.Printf("pattern is %s", r.URL.Path)
				http.NotFound(w, r) // 404
				return
			}
			fn(w, r)
		}
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError) // 500
		log.Println(err)
		return
	}
}
