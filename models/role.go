package models

import (
	"time"
)

type (
	// Role
	Role struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Name      string    `gorm:"not null;unique" json:"name"`
		CanWrite  *bool     `gorm:"default:true" json:"can_write"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      []User    `json:"-"`
	}
)
