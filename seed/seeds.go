package seed

import "gorm.io/gorm"

func All() []Seed {
	return []Seed{
		Seed{
			Name: "CreateSuperadmin",
			Run: func(db *gorm.DB) error {
				err := CreateRole(db, "Superadmin", true)
				return err
			},
		},
		Seed{
			Name: "CreateUser",
			Run: func(db *gorm.DB) error {
				err := CreateRole(db, "User", false)
				return err
			},
		},
	}
}
