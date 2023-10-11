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



## Features

- Search design patents by patent title, patent number, inventor(s) name, assignee (owner) name, application date, issue date, and design class.
- Efficiently parse and store USPTO design patent data.
- Optimize search engine performance for large datasets.

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
    

## Search Functionality
The Search engine uses fuzzy logic coupled with ElasticSearch (indexed against a postgres DB)
The search engine allows users to search for design patents based on various criteria, including patent title, patent number, inventor(s) name, assignee (owner) name, application date, issue date, and design class (if available). 
