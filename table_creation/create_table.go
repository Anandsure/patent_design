package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "anands"
	password  = "87szLCJM"
	dbname    = "design_patent"
	tableName = "patents"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table
	createTableStmt := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		PatentNumber VARCHAR(255) PRIMARY KEY,
		PatentTitle VARCHAR(255),
		Authors TEXT[],
		Assignee VARCHAR(255),
		ApplicationDate DATE,
		IssueDate DATE,
		DesignClass VARCHAR(255),
		ReferencesCited TEXT[],
		Description TEXT[]
	);`, tableName)

	_, err = db.Exec(createTableStmt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Table '%s' created successfully.\n", tableName)
	//inserting data into the table
	insert_data_into_table()
}
