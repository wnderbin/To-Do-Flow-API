package models

import (
	"time"

	"github.com/google/uuid"
)

type ToDoList struct {
	Id uuid.UUID `gorm:"primaryKey" json:"id"`

	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`

	ToDoNote string `json:"todo"`

	User_id int  `json:"user_id"`
	User    User `gorm:"foreignKey:User_id" json:"user"`
}
