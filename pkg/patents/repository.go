package patents

import (
	"github.com/Anandsure/patent_design/pkg/models"
)

type Repository interface {
	GetPatentJSONByNumber(patentNumber string) (*models.Patent, error)
}
