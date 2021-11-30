package models

import "time"

type (
	// Picture
	Picture struct {
		ID         uint      `json:"id" gorm:"primary_key"`
		PhoneID    uint      `json:"phoneID"`
		Title      string    `json:"title"`
		UrlPicture string    `gorm:"not null" json:"UrlPicture"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		Phone      Phone     `json:"-"`
	}
)
