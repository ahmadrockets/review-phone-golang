package seed

import (
	"final-project-sanber/models"

	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB, name string, canWrite bool) error {
	return db.Create(&models.Role{Name: name, CanWrite: &canWrite}).Error
}
