package models

import (
	"time"
)

type User struct {
	Id string `gorm:"type:text;primaryKey" json:"id"`

	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
