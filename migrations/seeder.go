package migrations

import (
	"github.com/Revprm/Nutrigrow-Backend/migrations/seeds" // Perbarui path import
	"gorm.io/gorm"
)

// Seeder menjalankan semua fungsi seeder yang terdaftar.
func Seeder(db *gorm.DB) error {
	// Panggil seeder user
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	// Panggil seeder kategori berita terlebih dahulu karena berita bergantung padanya
	if err := seeds.ListKategoriBeritaSeeder(db); err != nil {
		return err
	}

	// Panggil seeder berita
	if err := seeds.ListBeritaSeeder(db); err != nil {
		return err
	}

	// Panggil seeder bahan makanan
	if err := seeds.ListBahanMakananSeeder(db); err != nil {
		return err
	}

	// Panggil seeder makanan (bergantung pada bahan makanan)
	if err := seeds.ListMakananSeeder(db); err != nil {
		return err
	}

	// Panggil seeder stunting (bergantung pada user)
	if err := seeds.ListStuntingSeeder(db); err != nil {
		return err
	}

	return nil
}
