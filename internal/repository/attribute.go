package repository

import (
	"github.com/charanck/ABAC/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	CREATE_ATTRIBUTE_QUERY = "INSERT INTO attribute(id, name, string_value, integer_value, float_value, bool_value, date_value, type, updated, deleted, created) VALUES(?,?,?,?,?,?,?,?,?,?,?)"
)

type Attribute struct {
	db *sqlx.DB
}

func NewAttribute(db *sqlx.DB) Attribute {
	return Attribute{
		db: db,
	}
}

func (a *Attribute) Create(attribute model.Attribute) (string, error) {
	_, err := a.db.Exec(CREATE_ATTRIBUTE_QUERY, attribute.Id, attribute.Name, attribute.StringValue, attribute.IntegerValue, attribute.FloatValue, attribute.BoolValue, attribute.DateValue, attribute.Type, attribute.Updated, attribute.Deleted, attribute.Created)
	if err != nil {
		return "", nil
	}
	return attribute.Id, nil
}
