package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/public
var embeddedFiles embed.FS

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	fsys, err := fs.Sub(embeddedFiles, "frontend/public")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(fsys)))

	log.Printf("Starting server on %s", *addr)
	err = http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
