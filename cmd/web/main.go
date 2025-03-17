package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// fileServer := http.FileServer(http.Dir("./ui/html/base.tmpl.html/"))
	// mux.Handle("/base.tmpl.html/", http.StripPrefix("/base.tmpl.html", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/bug/view/", bugView)
	mux.HandleFunc("/bug/create/", bugCreate) 

	log.Println("Starting server on: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
