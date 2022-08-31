package main

import (
	_ "github.com/lib/pq"
	"hr-database-api/data"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := data.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
