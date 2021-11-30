package models

import "time"

type (
	// Review
	Review struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		PhoneID   uint      `json:"phoneID"`
		Title     string    `gorm:"not null" json:"title"`
		Content   string    `gorm:"not null" json:"content"`
		Picture   string    `json:"picture"`
		UserID    uint      `json:"userID"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Phone     Phone     `json:"-"`
		User      User      `json:"-"`
	}
)
