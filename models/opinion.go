package models

import "time"

type (
	// Opinion
	Opinion struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		PhoneID   uint      `json:"phoneID"`
		Nickname  string    `gorm:"not null" json:"nickname"`
		Content   string    `gorm:"not null" json:"content"`
		Picture   string    `json:"picture"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Phone     Phone     `json:"-"`
	}
)
