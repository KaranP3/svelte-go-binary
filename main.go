package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/public
var embeddedFiles embed.FS

func main() {
	fsys, err := fs.Sub(embeddedFiles, "frontend/public")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(":4000", nil)
}
