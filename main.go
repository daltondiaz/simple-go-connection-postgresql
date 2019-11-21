package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dalton00"
	dbname   = "test_go"
)

func main() {
	psqlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConnection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!!")

	GetResources(db)
}

// GetResources list all resources
func GetResources(db *sql.DB) {

	sql := "select id, name, description from resource"
	rows, err := db.Query(sql)

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var description string

		// the order of Scan variables its necessay
		// to be the same of query
		err = rows.Scan(&id, &name, &description)
		fmt.Println(id, name, description)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
