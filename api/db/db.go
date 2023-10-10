package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StringArray []string

// Value converts the string array to a string for database storage.
func (a StringArray) Value() (driver.Value, error) {
	return strings.Join(a, ","), nil
}

func (a *StringArray) Scan(value interface{}) error {
	str, ok := value.([]uint8)
	if !ok {
		return fmt.Errorf("failed to scan StringArray: unexpected value type %T", value)
	}

	var authors []string

	err := json.Unmarshal(str, &authors)
	if err != nil {
		return fmt.Errorf("failed to unmarshal StringArray: %v", err)
	}

	*a = authors
	return nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "anands"
	password = "87szLCJM"
	dbname   = "design_patent"
)

var db *gorm.DB = nil

// Patent represents the patent data structure
type Patent struct {
	PatentNumber    string      `gorm:"column:patentnumber;primaryKey"`
	PatentTitle     string      `gorm:"column:patenttitle"`
	Authors         StringArray `gorm:"column:authors"`
	Assignee        string      `gorm:"column:assignee"`
	ApplicationDate time.Time   `gorm:"column:applicationdate"`
	IssueDate       time.Time   `gorm:"column:issuedate"`
	DesignClass     string      `gorm:"column:designclass"`
	ReferencesCited StringArray `gorm:"column:referencescited"`
	Description     StringArray `gorm:"column:description"`
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	db = Connect()
	return db
}

func Connect() *gorm.DB {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable password=%s", host, user, dbname, port, password)
	sqlDB, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormConfig := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), gormConfig)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetPatentJSONByNumber(patentNumber string) (string, error) {
	var patent Patent

	if err := db.Model(&Patent{}).Where("patentnumber = ?", patentNumber).First(&patent).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		log.Printf("Error fetching patent by number: %v", err)
		return "", err
	}

	// Convert patent to JSON
	patentJSON, err := json.Marshal(patent)
	if err != nil {
		log.Printf("Error converting patent to JSON: %v", err)
		return "", err
	}

	return string(patentJSON), nil
}
