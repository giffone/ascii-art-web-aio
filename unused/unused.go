package unused

// func main() {
// 	http.HandleFunc("/", handlers.Home(memory))
// }

// http.HandleFunc("/", handlers.CheckHandler(handler.Home))
// http.HandleFunc("/", handlers.Home(memory))
// http.HandleFunc("/ascii-art", handlers.CheckHandler(handler.AsciiArt))
// http.HandleFunc("/download", handlers.CheckHandler(handler.Download))

// func Home(p *structure.Art) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == "GET" {
// 			tmpl.ExecuteTemplate(w, "index", p.Page)
// 		}
// 	}
// }

// func (h *HandlerArt) Home(w http.ResponseWriter, r *http.Request, ptrn string) {
// 	if r.Method == "POST" {
// 		h.Art.Page.Clear()
// 	}
// 	renderTemplate(w, "index", &h.Art.Page)
// }

// func (h *HandlerArt) AsciiArt(w http.ResponseWriter, r *http.Request, ptrn string) {
// 	if r.Method == "POST" {
// 		b := new(structure.Buf)
// 		h.Art.Page.Add(r, "submit")
// 		art.ArtMain(h.Art, b)
// 		renderTemplate(w, "index", &h.Art.Page)
// 	}
// }

// func (h *HandlerArt) Download(w http.ResponseWriter, r *http.Request, ptrn string) {
// 	if r.Method == "POST" {
// 		h.Art.Page.Add(r, "save")
// 		if filename := output.Output(h.Art); filename != "" {
// 			buf := &bytes.Buffer{}
// 			for _, line := range h.Art.Page.AnsFile {
// 				buf.WriteString(line)
// 				buf.WriteRune('\n')
// 			}

// 			file := bytes.NewReader(buf.Bytes())
// 			w.Header().Set("Content-Type", "text/plain")
// 			w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
// 			w.Header().Set("Content-Length", fmt.Sprintf("%d", file.Size()))
// 			io.Copy(w, file)
// 		}
// 		renderTemplate(w, "index", &h.Art.Page)
// 	}
// }





// func (h *HandlerArt) MainHandler(w http.ResponseWriter, r *http.Request, ptrn string) {
// 	if r.Method == "POST" {
// 		if ptrn == "/" {
// 			h.Art.Page.Clear()
// 			renderTemplate(w, "index", &h.Art.Page)
// 			log.Println(logging.FormatRequest(r))
// 		} else if ptrn == "ascii-art" {
// 			b := new(structure.Buf)

// 			h.Art.Page.Add(r, "submit")
// 			art.ArtMain(h.Art, b)
// 			renderTemplate(w, "ifunc (h *HandlerArt) MainHandler(w http.ResponseWriter, r *http.Request, ptrn string) {
// 	if r.Method == "POST" {
// 		if ptrn == "/" {
// 			h.Art.Page.Clear()
// 			renderTemplate(w, "index", &h.Art.Page)
// 			log.Println(logging.FormatRequest(r))
// 		} else if ptrn == "ascii-art" {
// 			b := new(structure.Buf)

// 			h.Art.Page.Add(r, "submit")
// 			art.ArtMain(h.Art, b)
// 			renderTemplate(w, "index", &h.Art.Page)
// 			log.Println(logging.FormatRequest(r))
// 		} else if ptrn == "download" {
// 			h.Art.Page.Add(r, "save")

// 			filename := output.Output(h.Art)
// 			arr := []byte{}

// 			for _, line := range h.Art.Page.AnsFile {
// 				arr = append(arr, []byte(line)...)
// 				arr = append(arr, '\n')
// 			}
// 			w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
// 			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(arr)))

// 			if _, err := w.Write(arr); err != nil {
// 				logging.WarningErr(h.Art, err, err.Error())
// 				return
// 			}
// 			log.Println(logging.FormatRequest(r))
// 		}
// 	} else if r.Method == "GET" {
// 		if ptrn == "/" {
// 			renderTemplate(w, "index", &h.Art.Page)
// 			log.Println(logging.FormatRequest(r))
// 		} else if ptrn == "favicon.ico" {
// 			w.Header().Set("Content-Type", "image/x-icon")
// 			w.Header().Set("Cache-Control", "public, max-age=7776000")
// 			http.ServeFile(w, r, "ico/favicon.ico")
// 		} else {
// 			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed) // 405
// 			log.Println(logging.FormatRequest(r))
// 			return
// 		}
// 	}
// }
// 			filename := output.Output(h.Art)
// 			arr := []byte{}

// 			for _, line := range h.Art.Page.AnsFile {
// 				arr = append(arr, []byte(line)...)
// 				arr = append(arr, '\n')
// 			}
// 			w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
// 			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(arr)))

// 			if _, err := w.Write(arr); err != nil {
// 				logging.WarningErr(h.Art, err, err.Error())
// 				return
// 			}
// 			log.Println(logging.FormatRequest(r))
// 		}
// 	} else if r.Method == "GET" {
// 		if ptrn == "/" {
// 			renderTemplate(w, "index", &h.Art.Page)
// 			log.Println(logging.FormatRequest(r))
// 		} else if ptrn == "favicon.ico" {
// 			w.Header().Set("Content-Type", "image/x-icon")
// 			w.Header().Set("Cache-Control", "public, max-age=7776000")
// 			http.ServeFile(w, r, "ico/favicon.ico")
// 		} else {
// 			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed) // 405
// 			log.Println(logging.FormatRequest(r))
// 			return
// 		}
// 	}
// }