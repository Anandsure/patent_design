package bulk_insertion

import (
	"github.com/Anandsure/patent_design/api/utils"
)

func StartBulk() {
	utils.ImportEnv()

	// gormDb := db.GetDB()
	// gormDb.AutoMigrate(&models.Patent{})
	// Inserting data into the table
	insertData()
	esBulkInsert()
}
