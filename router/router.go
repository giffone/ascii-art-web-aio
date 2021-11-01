package router

import (
	"net/http"
)

func NewRouter(h *HandlerArt) *http.ServeMux {
	r := http.NewServeMux()
	dir := http.Dir("assets")
	dirHandler := http.StripPrefix("/assets/", http.FileServer(dir))

	r.Handle("/assets/", dirHandler)
	r.HandleFunc("/", checkHandler(h.mainHandler))
	r.HandleFunc("/ascii-art", checkHandler(h.artHandler))
	r.HandleFunc("/dowload", checkHandler(h.downloadHandler))

	return r
}
