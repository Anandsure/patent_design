package users

import (
	"github.com/Anandsure/patent_design/pkg/models"
	"github.com/google/uuid"
)

type Repository interface {
	Find(id *uuid.UUID) (*models.Users, error)
}
