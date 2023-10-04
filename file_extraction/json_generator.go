package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type UsPatentGrant struct {
	XMLName              xml.Name   `xml:"us-patent-grant"`
	PublicationReference DocumentID `xml:"us-bibliographic-data-grant>publication-reference>document-id"`
	ApplicationReference DocumentID `xml:"us-bibliographic-data-grant>application-reference>document-id"`
	InventionTitle       string     `xml:"us-bibliographic-data-grant>invention-title"`
}

type DocumentID struct {
	Country   string `xml:"country"`
	DocNumber string `xml:"doc-number"`
	Kind      string `xml:"kind"`
	Date      string `xml:"date"`
}

func main() {
	xmlFilePath := "all_xml/USD0984780-20230502.XML"

	xmlData, err := os.ReadFile(xmlFilePath)
	if err != nil {
		fmt.Printf("Error reading XML file: %v", err)
		return
	}

	var patentGrant UsPatentGrant
	err = xml.Unmarshal(xmlData, &patentGrant)
	if err != nil {
		fmt.Printf("Error unmarshaling XML: %v", err)
		return
	}

	// Convert the structure to JSON
	jsonData, err := json.MarshalIndent(patentGrant, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	// Save the JSON data to a file
	jsonFilePath := "patent_data.json"
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON file: %v", err)
		return
	}

	fmt.Println("JSON file created successfully.")
}
