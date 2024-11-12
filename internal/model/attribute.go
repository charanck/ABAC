package model

import "time"

type Attribute struct {
	Id           string    `db:"id"`
	Name         string    `db:"name"`
	StringValue  string    `db:"string_value"`
	IntegerValue int       `db:"integer_value"`
	FloatValue   float32   `db:"float_value"`
	BoolValue    bool      `db:"bool_value"`
	DateValue    time.Time `db:"date_value"`
	Type         string    `db:"type"`
	Updated      time.Time `db:"updated"`
	Deleted      time.Time `db:"deleted"`
	Created      time.Time `db:"created"`
}

func (r Attribute) GetTableName() string {
	return "attribute"
}
