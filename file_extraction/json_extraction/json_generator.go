package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Inventor struct {
	LastName  string `xml:"addressbook>last-name"`
	FirstName string `xml:"addressbook>first-name"`
}

type UsPatentGrant struct {
	PatentTitle     string      `xml:"us-bibliographic-data-grant>invention-title"`
	PatentNumber    string      `xml:"us-bibliographic-data-grant>publication-reference>document-id>doc-number"`
	Authors         []Inventor  `xml:"us-bibliographic-data-grant>us-parties>inventors>inventor"`
	Assignee        string      `xml:"us-bibliographic-data-grant>us-parties>us-applicants>us-applicant>addressbook>orgname"`
	ApplicationDate CustomTime  `xml:"us-bibliographic-data-grant>application-reference>document-id>date"`
	IssueDate       CustomTime  `xml:"us-bibliographic-data-grant>publication-reference>document-id>date"`
	DesignClass     string      `xml:"us-bibliographic-data-grant>classification-national>main-classification"`
	ReferencesCited []Reference `xml:"us-bibliographic-data-grant>us-references-cited>us-citation,omitempty"`
	Description     Description `xml:"description"`
}

type Reference struct {
	Name string `xml:"patcit>document-id>name"`
}

type Description struct {
	DescriptionDrawings []string `xml:"description-of-drawings>p"`
}

type CustomTime struct {
	Time string `xml:",chardata"`
}

func processXMLFile(xmlFilePath string) (map[string]interface{}, error) {
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening XML file %s: %v", xmlFilePath, err)
	}
	defer xmlFile.Close()

	var patent UsPatentGrant

	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&patent); err != nil {
		return nil, fmt.Errorf("error decoding XML for file %s: %v", xmlFilePath, err)
	}

	// Process description
	descriptionDrawings := make([]string, len(patent.Description.DescriptionDrawings))
	for i, desc := range patent.Description.DescriptionDrawings {
		// Remove newline characters and spaces
		desc = strings.TrimSpace(strings.ReplaceAll(desc, "\n", ""))
		descriptionDrawings[i] = desc
	}

	// Process references
	var references []string
	for _, ref := range patent.ReferencesCited {
		if ref.Name != "" {
			references = append(references, ref.Name)
		}
	}

	// Process authors
	var authors []string
	for _, inventor := range patent.Authors {
		authorName := fmt.Sprintf("%s %s", inventor.FirstName, inventor.LastName)
		authors = append(authors, authorName)
	}

	return map[string]interface{}{
		"PatentTitle":     patent.PatentTitle,
		"PatentNumber":    patent.PatentNumber,
		"Authors":         authors,
		"Assignee":        patent.Assignee,
		"ApplicationDate": patent.ApplicationDate.Time,
		"IssueDate":       patent.IssueDate.Time,
		"DesignClass":     patent.DesignClass,
		"ReferencesCited": references,
		"Description":     descriptionDrawings,
	}, nil
}

func main() {
	xmlFiles, err := filepath.Glob("../all_xml/*.XML")
	if err != nil {
		fmt.Printf("Error finding XML files: %v\n", err)
		return
	}

	var combinedData []map[string]interface{}

	for _, xmlFilePath := range xmlFiles {
		data, err := processXMLFile(xmlFilePath)
		if err != nil {
			fmt.Printf("Error processing XML file %s: %v\n", xmlFilePath, err)
			continue
		}

		combinedData = append(combinedData, data)
	}

	jsonData, err := json.MarshalIndent(combinedData, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}

	outputFile, err := os.Create("combined_patent_data.json")
	if err != nil {
		fmt.Printf("Error creating JSON file: %v\n", err)
		return
	}
	defer outputFile.Close()

	_, err = io.WriteString(outputFile, string(jsonData))
	if err != nil {
		fmt.Printf("Error writing JSON to file: %v\n", err)
		return
	}

	fmt.Println("Combined JSON data written to combined_patent_data.json")
}
