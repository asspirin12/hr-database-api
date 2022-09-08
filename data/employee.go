package data

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 54333
	user     = "guest"
	password = "guest"
	dbname   = "guest"
)

type Employee struct {
	Id int `json:"id"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Department string `json:"department"`

	Email string `json:"email,omitempty"`

	DateHired string `json:"date_hired"`
}

// ConnectDatabase checks the connection to the database
func ConnectDatabase() error {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to database")
	DB = db
	return nil
}

// GetEmployees returns a list of employees
func GetEmployees() ([]Employee, error) {
	statement := `select id, first_name, last_name, email, department, date_hired from employees limit 10`

	rows, err := DB.Query(statement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	workforce := make([]Employee, 0)

	for rows.Next() {
		person := Employee{}
		err = rows.Scan(
			&person.Id,
			&person.FirstName,
			&person.LastName,
			&person.Email,
			&person.Department,
			&person.DateHired,
		)
		if err != nil {
			return nil, err
		}

		workforce = append(workforce, person)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return workforce, nil
}
