package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

	Created_at time.Time `json:"created_at"`
}
