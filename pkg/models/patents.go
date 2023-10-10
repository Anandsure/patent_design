package models

import (
	"github.com/lib/pq"
)

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
