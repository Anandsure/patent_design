package migrations

import (
	"github.com/Anandsure/patent_design/api/db"
	"github.com/Anandsure/patent_design/pkg/models"
)

func Migrate() {
	database := db.GetDB()
	database.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	database.AutoMigrate(&models.Patent{})
}
