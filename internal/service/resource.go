package service

import (
	"time"

	"github.com/charanck/ABAC/internal/model"
	"github.com/charanck/ABAC/internal/repository"
	"github.com/charanck/ABAC/internal/util"
	"github.com/google/uuid"
)

type Resource struct {
	repository *repository.Resource
}

func NewResource(repository *repository.Resource) Resource {
	return Resource{
		repository: repository,
	}
}

func (r *Resource) Create(resource model.Resource) (string, error) {
	resource.Id = uuid.New().String()
	resource.Created = time.Now()
	resource.Updated = time.Now()
	existingResource, err := r.getByName(resource.Name)
	if err != nil {
		return "", err
	}
	if existingResource.Id != "" {
		return "", util.ErrAlreadyExists(nil, "Resource with the name already exists")
	}
	// Before creating a resource verify if the policy exists
	return r.repository.Create(resource)
}

func (r *Resource) getByName(resourceName string) (model.Resource, error) {
	return r.repository.GetByName(resourceName)
}

func (r *Resource) GetById(resourceId string) (model.Resource, error) {
	return r.repository.GetById(resourceId)
}

func (r *Resource) List() ([]model.Resource, error) {
	return r.repository.List()
}

func (r *Resource) DeleteById(resourceId string) (string, error) {
	return r.repository.DeleteById(resourceId)
}

func (r *Resource) UpdateById(resource model.Resource, fieldMask []string) (string, error) {
	resource.Updated = time.Now()
	fieldMask = append(fieldMask, "updated")
	r.repository.Update(resource, fieldMask)
	return "", nil
}
