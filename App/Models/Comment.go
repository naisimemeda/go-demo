package Models

import "go-demo/App/Api"

type Comment struct {
	ID        uint         `gorm:"primary_key" json:"id"`
	Content   string       `json:"content"`
	PostID    string       `json:"post_id"`
	UserID    string       `json:"user_id"`
	CreatedAt Api.JSONTime `json:"created_at"`
}
