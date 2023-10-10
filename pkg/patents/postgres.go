package patents

import (
	"errors"

	"github.com/Anandsure/patent_design/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) GetPatentJSONByNumber(patentNumber string) (*models.Patent, error) {
	var patent models.Patent
	if err := r.DB.Model(&models.Patent{}).Where("patent_number = ?", patentNumber).First(&patent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &patent, nil
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
