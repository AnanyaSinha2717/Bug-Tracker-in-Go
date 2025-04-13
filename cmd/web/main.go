package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errlog  *log.Logger
	infolog *log.Logger
}

func main() {
	// cmd line flag, completely optional
	addr := flag.String("addr", ":4000", "HTTP Network Address") // this is a pointer which we deref  later
	dsn := flag.String("dsn", "lily:BOSS@/buggo?parseTime=true", "MySQL data source name")

	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err:= openDB(*dsn)
	if err != nil {
		errlog.Fatal(err)
	}

	defer db.Close()

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
	err = srv.ListenAndServe()
	errlog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := mysql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}