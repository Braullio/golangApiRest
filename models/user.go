package models

import "time"

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Status  string `json:"status"`
	Created string `json:"_"`
	Updated string `json:"_"`

	CreatedTime time.Time `json:"created_at"`
	UpdatedTime time.Time `json:"updated_at"`
}
