package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// cmd line flag, completely optional
	addr := flag.String("addr", ":4000", "HTTP Network Address") // this is a pointer which we deref  later

	flag.Parse()

	mux := http.NewServeMux()
	// fileServer := http.FileServer(http.Dir("./ui/html/base.tmpl.html/"))
	// mux.Handle("/base.tmpl.html/", http.StripPrefix("/base.tmpl.html", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/bug/view/", bugView)
	mux.HandleFunc("/bug/create/", bugCreate) 

	log.Printf("Starting server on: %s", *addr)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
