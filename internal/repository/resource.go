package repository

import (
	"github.com/charanck/ABAC/internal/model"
	querybuilder "github.com/charanck/ABAC/internal/model/query_builder"
	"github.com/jmoiron/sqlx"
)

const (
	CREATE_RESOURCE_QUERY = "INSERT INTO resource (id, name, owner_id, policy_id, description, updated, deleted, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	GET_RESOURCE_BY_ID    = "SELECT * FROM resource WHERE id = ?"
	GET_RESOURCE_BY_NAME  = "SELECT * FROM resource WHERE name = ?"
	LIST_RESOURCE         = "SELECT * FROM resource LIMIT ? OFFSET ?"
	DELETE_RESOURCE_BY_ID = "DELETE FROM resource WHERE id = ?"
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

func (r *Resource) List(limit, offset int) ([]model.Resource, error) {
	rows, err := r.db.Queryx(LIST_RESOURCE, limit, offset)
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

func (r *Resource) DeleteById(resourceId string) (string, error) {
	_, err := r.db.Exec(DELETE_RESOURCE_BY_ID, resourceId)
	if err != nil {
		return "", err
	}
	return resourceId, nil
}

func (r *Resource) Update(resource model.Resource, fieldMask []string) (string, error) {
	query, queryValues, err := querybuilder.BuildUpdateQuery(&resource, fieldMask, querybuilder.Where{
		Left: &querybuilder.Where{
			Value: "id",
		},
		Right: &querybuilder.Where{
			Value: resource.Id,
		},
		Operation: querybuilder.Equal{},
	})
	if err != nil {
		return "", err
	}
	_, err = r.db.Exec(query, queryValues...)
	if err != nil {
		return "", err
	}
	return resource.Id, nil
}
