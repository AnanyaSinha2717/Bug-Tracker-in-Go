package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errlog  *log.Logger
	infolog *log.Logger
}

func main() {
	// cmd line flag, completely optional
	addr := flag.String("addr", ":4000", "HTTP Network Address") // this is a pointer which we deref  later

	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errlog: errlog,
		infolog:  infolog,
	}


	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errlog,
		Handler:  app.routes(),
	}

	infolog.Printf("Starting server on: %s", *addr)
	err := srv.ListenAndServe()
	errlog.Fatal(err)
}
