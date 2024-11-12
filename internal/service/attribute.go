package service

import (
	"time"

	"github.com/charanck/ABAC/internal/model"
	"github.com/charanck/ABAC/internal/repository"
	"github.com/google/uuid"
)

type Attribute struct {
	repository *repository.Attribute
}

func NewAttribute(repository *repository.Attribute) Attribute {
	return Attribute{
		repository: repository,
	}
}

func (a *Attribute) Create(attribute model.Attribute) (string, error) {
	attribute.Id = uuid.New().String()
	attribute.Created = time.Now()
	attribute.Updated = time.Now()
	return a.repository.Create(attribute)
}
