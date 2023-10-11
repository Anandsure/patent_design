package bulk_insertion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Anandsure/patent_design/api/db"
	"github.com/Anandsure/patent_design/pkg/models"
)

const (
	// Little automation left due to lack of time
	jsonFile = "../file_extraction/json_extraction/combined_patent_data.json"
)

func insertData() {
	// Read the JSON data
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var patentData []models.Patent
	err = json.Unmarshal(byteValue, &patentData)
	if err != nil {
		log.Fatal(err)
	}

	// Establish a PostgreSQL connection
	gormDB := db.GetDB()

	// Insert data into the table
	for _, data := range patentData {
		t, err := time.Parse("20060102", data.ApplicationDate)
		if err != nil {
			continue
		}
		x, err := time.Parse("20060102", data.IssueDate)
		if err != nil {
			continue
		}

		data.ApplicationDate = fmt.Sprintf("%d", t.Unix())
		data.IssueDate = fmt.Sprintf("%d", x.Unix())

		err = gormDB.Create(&data).Error
		if err != nil {
			log.Printf("Error inserting data for patent %s: %v", data.PatentNumber, err)
		} else {
			fmt.Printf("Data inserted successfully for patent %s!\n", data.PatentNumber)
		}
	}
}
