package migrations

import (
	"github.com/Revprm/Nutrigrow-Backend/entity" // Ensure this path is correct
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{}, // Existing
		&entity.RefreshToken{}, // Existing

		&entity.Stunting{},
		&entity.KategoriBerita{},
		&entity.Berita{},
		&entity.BahanMakanan{},
		&entity.Makanan{},
	); err != nil {
		return err
	}

	return nil
}