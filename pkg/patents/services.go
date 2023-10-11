package patents

import (
	"github.com/Anandsure/patent_design/pkg/models"
)

type Service interface {
	GetPatentJSONByNumber(patentNumber string) (*models.Patent, error)
}

type userSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}

func (u *userSvc) GetPatentJSONByNumber(patentNumber string) (*models.Patent, error) {
	return u.repo.GetPatentJSONByNumber(patentNumber)
}
