package db

import "github.com/Anandsure/patent_design/pkg/patents"

var (
	PatentSvc patents.Service = nil
)

func InitServices() {
	db := GetDB()

	patentsRepo := patents.NewPostgresRepo(db)
	PatentSvc = patents.NewService(patentsRepo)
}
