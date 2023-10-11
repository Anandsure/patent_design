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

   ```

### Usage

    go mod download
    go run main.go

## Search Functionality

The Search engine uses fuzzy logic coupled with ElasticSearch (indexed against a postgres DB)
The search engine allows users to search for design patents based on various criteria, including patent title, patent number, inventor(s) name, assignee (owner) name, application date, issue date, and design class (if available).
