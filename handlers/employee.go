package handlers

import (
	"encoding/json"
	"hr-database-api/data"
	"log"
	"net/http"
)

type Employees struct{}

func (e Employees) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/employees" {
		if r.Method == http.MethodGet {
			//rw.Write([]byte(("This is hr-database-api")))
			e.getEmployees(rw, r)
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

// getEmployees returns a list of employees (limit is 10 records by default)
func (e Employees) getEmployees(rw http.ResponseWriter, r *http.Request) {
	log.Println("Handle GET employees")

	employeesList, err := data.GetEmployees()
	if err != nil {
		http.Error(rw, "Unable to get employees list", http.StatusInternalServerError)
		log.Println(err)
	}

	encoder := json.NewEncoder(rw)

	err = encoder.Encode(employeesList)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
		log.Println(err)
	}
}
