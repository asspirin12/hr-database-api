package main

import (
	_ "github.com/lib/pq"
	"hr-database-api/data"
	"hr-database-api/handlers"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := data.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	eh := handlers.Employees{}
	mux := http.NewServeMux()

	mux.Handle("/employees", eh)

	s := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
