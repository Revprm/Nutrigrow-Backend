package seeds

import (
	"errors"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

// ListKategoriBeritaSeeder menyemai data kategori berita awal ke dalam database.
func ListKategoriBeritaSeeder(db *gorm.DB) error {
	kategoriBeritaList := []entity.KategoriBerita{
		{NamaKategoriBerita: "Gizi Anak"},
		{NamaKategoriBerita: "Kesehatan Ibu"},
		{NamaKategoriBerita: "Resep Makanan"},
		{NamaKategoriBerita: "Pencegahan Stunting"},
	}

	// Memastikan tabel ada sebelum menyemai data
	hasTable := db.Migrator().HasTable(&entity.KategoriBerita{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.KategoriBerita{}); err != nil {
			return err
		}
	}

	for _, data := range kategoriBeritaList {
		var kategori entity.KategoriBerita
		// Memeriksa apakah data kategori berita sudah ada berdasarkan nama
		err := db.Where(&entity.KategoriBerita{NamaKategoriBerita: data.NamaKategoriBerita}).First(&kategori).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// Jika data belum ada, buat record baru
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
