package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "anands"
	password  = "87szLCJM"
	dbname    = "design_patent"
	tableName = "patents"
)

var db *gorm.DB = nil

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	db = Connect()
	return db
}

func Connect() *gorm.DB {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable password=%s", host, user, dbname, port, password)
	sqlDB, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
