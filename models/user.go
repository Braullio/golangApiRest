package models

import (
	"cloud.google.com/go/civil"
)

type User struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Phone   string         `json:"phone"`
	Status  string         `json:"status"`
	Created civil.DateTime `json:"created_at"`
	Updated civil.DateTime `json:"updated_at"`
}
