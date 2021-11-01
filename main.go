package main

import (
	"ascii-art-web/art/scan"
	"ascii-art-web/router"
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func main() {
	// port := os.Getenv("PORT")
	memory := router.NewHandlerArt()
	r := router.NewRouter(memory)

	// adding fonts
	scan.Directory(memory.Art, "art/fonts")
	// adding colors
	memory.Art.Page.AddColors()
	// adding formats
	memory.Art.Page.AddFormats()

	log.Println("server is running...")
	fmt.Println("Link http://localhost:8080/")
	log.Println(http.ListenAndServe(port, r))
}
