package Models

import "go-demo/App/Api"

type Post struct {
	ID        uint         `gorm:"primary_key" json:"id"`
	Title     string       `json:"title"`
	UserID    string       `json:"user_id"`
	CreatedAt Api.JSONTime `json:"created_at"`
}
