package model

import "time"

type Resource struct {
	Id          string
	Name        string
	OwnerId     string
	PolicyId    string
	Description string
	Updated     time.Time
	Deleted     time.Time
	Created     time.Time
}
