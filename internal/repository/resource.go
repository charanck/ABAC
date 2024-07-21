package repository

import (
	"github.com/charanck/ABAC/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	CREATE_RESOURCE_QUERY = "INSERT INTO resource (id, name, owner_id, policy_id, description, updated, deleted, created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	GET_RESOURCE_BY_ID    = "SELECT * FROM resource WHERE id = $1"
	GET_RESOURCE_BY_NAME  = "SELECT * FROM resource WHERE name = $1"
	LIST_RESOURCE         = "SELECT * FROM resource"
)

type Resource struct {
	db *sqlx.DB
}

func NewResource(db *sqlx.DB) Resource {
	return Resource{
		db: db,
	}
}

func (r *Resource) Create(resource model.Resource) (string, error) {
	_, err := r.db.Exec(CREATE_RESOURCE_QUERY, resource.Id, resource.Name, resource.OwnerId, resource.PolicyId, resource.Description, resource.Updated, resource.Deleted, resource.Created)
	if err != nil {
		return "", err
	}
	return resource.Id, nil
}

func (r *Resource) GetById(resourceId string) (model.Resource, error) {
	rows, err := r.db.Queryx(GET_RESOURCE_BY_ID, resourceId)
	if err != nil {
		return model.Resource{}, err
	}
	resource := model.Resource{}
	for rows.Next() {
		rows.StructScan(&resource)
	}
	return resource, nil
}

func (r *Resource) GetByName(resourceName string) (model.Resource, error) {
	rows, err := r.db.Queryx(GET_RESOURCE_BY_NAME, resourceName)
	if err != nil {
		return model.Resource{}, err
	}
	resource := model.Resource{}
	for rows.Next() {
		err := rows.StructScan(&resource)
		if err != nil {
			return model.Resource{}, err
		}
	}
	return resource, nil
}

func (r *Resource) ListResource() ([]model.Resource, error) {
	rows, err := r.db.Queryx(LIST_RESOURCE)
	if err != nil {
		return nil, err
	}
	resources := []model.Resource{}
	for rows.Next() {
		currentResource := model.Resource{}
		err := rows.StructScan(&currentResource)
		if err != nil {
			return nil, err
		}
		resources = append(resources, currentResource)
	}
	return resources, nil
}
