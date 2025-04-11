package models

import (
	"time"
)

type ToDoList struct {
	Id string `gorm:"primaryKey" json:"id"`

	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`

	Todonote string `json:"todo"`

	User_id string `json:"user_id"`
	User    User   `gorm:"foreignKey:User_id" json:"user"`
}
