package Models

import "go-demo/App/Api"

type User struct {
	ID        uint         `gorm:"primary_key" json:"id"`
	Name      string       `json:"name"`
	CreatedAt Api.JSONTime `json:"created_at"`
	Post      []Post
}
