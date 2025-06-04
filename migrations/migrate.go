package migrations

import (
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.RefreshToken{},
	); err != nil {
		return err
	}

	return nil
}
