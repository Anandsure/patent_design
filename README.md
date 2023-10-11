# USPTO Design Patent Search Engine

This project implements a search engine for USPTO design patents based on various criteria. Users can search for design patents by patent title, patent number, inventor(s) name, assignee (owner) name, application date, issue date, and design class (if available).

## Table of Contents

- [Introduction](#introduction)
- [Search Engine Architecture](#data-diagram)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Database](#database)
- [Search Functionality](#search-functionality)
- [Performance Optimization](#performance-optimization)
- [Documentation](#documentation)
- [License](#license)

## Introduction

The United States Patent and Trademark Office (USPTO) provides a dataset of design patents, including information about various design patents granted by the USPTO. This project aims to create a search engine that enables users to search for design patents based on specific criteria.

## Search Engine Architecture
### Flow 1 -> Json Extraction and bulk inserts into postgreSQL and elasticSearch
<img width="705" alt="image" src="https://github.com/Anandsure/patent_design/assets/43916800/fb5e0c09-6e4e-44b3-b8cc-99044d486b26">

### Flow 2 -> Optimised wildcard search from ElasticSearch with all hits and Pagination
<img width="837" alt="image" src="https://github.com/Anandsure/patent_design/assets/43916800/a321bdd8-d7ab-4215-a5b1-dfd95e02a75b">

### Flow 3 -> Optimised Query to get all the patent metadata from PostgreSQL 
<img width="828" alt="image" src="https://github.com/Anandsure/patent_design/assets/43916800/0d3c8932-8514-478b-ba1c-6d8baac89808">


## Features

- Search design patents by patent title, patent number, inventor(s) name, assignee (owner) name, application date, issue date, and design class.
- Efficiently parse and store USPTO design patent data.
- Optimize search engine performance for large datasets.

## Patent Model
```
type Patent struct {
	PatentNumber    string         `json:"PatentNumber" gorm:"primaryKey"`
	PatentTitle     string         `json:"PatentTitle"`
	Authors         pq.StringArray `json:"Authors" gorm:"type:text[]"`
	Assignee        string         `json:"Assignee"`
	ApplicationDate string         `json:"ApplicationDate"`
	IssueDate       string         `json:"IssueDate"`
	DesignClass     string         `json:"DesignClass"`
	ReferencesCited pq.StringArray `json:"ReferencesCited" gorm:"type:text[]"`
	Description     pq.StringArray `json:"Description" gorm:"type:text[]"`
}
```
The fields have been highly optimiised to hold list of data, the extraction has been done refering to the dtd from the USPTO page for design patents

## Getting Started

### Prerequisites

To run this project, you need the following prerequisites:

- GoLang (v1.20)
- PostgreSQL (v12+)
- ElasticSearch (v17.17)

### Installation

1. Clone the repository:
    ```bash
   git clone https://github.com/yourusername/patent_designs.git
   cd patent_designs

### Usage
    go mod download
    go run main.go

### Docker Usage - Make sure you have docker-compose installed
    docker-compose up -d //runs in detached mode as a daemon 

## Search Functionality
The Search engine uses fuzzy logic coupled with ElasticSearch (indexed against a postgres DB)
The search engine allows users to search for design patents based on various criteria, including patent title, patent number, inventor(s) name, assignee (owner) name, application date, issue date, and design class (if available). 

## Performance Optimisation
according to the 2 Flows,
1. The search engine data in ES only has the searchable fields (title, authors, asignee, etc;) along with the unique key (patent number)
the returned response is a linear json that can be used in the front-end as a list view , to show the search results immediately.  <br>
2. The ES results are paginated , making sure it doesn't take a hit on the performance. <br>
3. when the User wants to retrieve all the data about the patent, the data is queried with the patent_number(pk) field from the postgres DB directly.
The returned response is now a full metadata json with all the fields pertaining to the Patent. <br>
4. The Search query uses the wildcard mode to effectively apply logic on all requested fields and weigh it by the most likey score. 

## Documentation
- [Postman Documentation](https://documenter.getpostman.com/view/9325142/2s9YJjQden) Please refer to this for API documentation. 
