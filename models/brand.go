package models

type (
	// Brand
	Brand struct {
		ID    uint    `json:"id" gorm:"primary_key"`
		Name  string  `gorm:"not null;unique" json:"name"`
		Phone []Phone `json:"-"`
	}
)
