package model

import "time"

const tableName string = "resource"

type Resource struct {
	Id          string    `db:"id"`
	Name        string    `db:"name"`
	OwnerId     string    `db:"owner_id"`
	PolicyId    string    `db:"policy_id"`
	Description string    `db:"description"`
	Updated     time.Time `db:"updated"`
	Deleted     time.Time `db:"deleted"`
	Created     time.Time `db:"created"`
}

func (r Resource) GetTableName() string {
	return tableName
}
