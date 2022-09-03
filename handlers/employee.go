package handlers

import (
	"log"
	"net/http"
)

type Employees struct{}

func (e Employees) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/employees" {
		if r.Method == http.MethodGet {
			rw.Write([]byte(("This is hr-database-api")))
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			log.Printf("Method %s is not allowed for %s", r.Method, r.URL.Path)
			return
		}
	} else {
		rw.WriteHeader(http.StatusNotFound)
		log.Fatalf("Page %s is not found", r.URL.Path)
		return
	}
}
