package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbName   = "design_patent"
	jsonFile = "../file_extraction/json_extraction/combined_patent_data.json"
)

type PatentData struct {
	PatentNumber    string   `json:"PatentNumber"`
	PatentTitle     string   `json:"PatentTitle"`
	Authors         []string `json:"Authors"`
	Assignee        string   `json:"Assignee"`
	ApplicationDate string   `json:"ApplicationDate"`
	IssueDate       string   `json:"IssueDate"`
	DesignClass     string   `json:"DesignClass"`
	ReferencesCited []string `json:"ReferencesCited"`
	Description     []string `json:"Description"`
}

func insert_data_into_table() {
	// Read the JSON data
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var patentData []PatentData
	err = json.Unmarshal(byteValue, &patentData)
	if err != nil {
		log.Fatal(err)
	}

	// Establish a PostgreSQL connection
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert data into the table
	for _, data := range patentData {
		// Convert slices to JSON strings
		authorsJSON, err := json.Marshal(data.Authors)
		if err != nil {
			log.Fatal(err)
		}

		referencesJSON, err := json.Marshal(data.ReferencesCited)
		if err != nil {
			log.Fatal(err)
		}

		descriptionJSON, err := json.Marshal(data.Description)
		if err != nil {
			log.Fatal(err)
		}

		sqlStatement := fmt.Sprintf(`
			INSERT INTO %s (PatentNumber, PatentTitle, Authors, Assignee, ApplicationDate, IssueDate, DesignClass, ReferencesCited, Description)
			VALUES ($1, $2, $3::jsonb, $4, $5, $6, $7, $8::jsonb, $9::jsonb)`, tableName)

		_, err = db.Exec(sqlStatement, data.PatentNumber, data.PatentTitle, authorsJSON, data.Assignee, data.ApplicationDate, data.IssueDate, data.DesignClass, referencesJSON, descriptionJSON)
		if err != nil {
			log.Fatalf("Error inserting data: %v", err)
		}
	}

	fmt.Println("Data inserted successfully!")
}
