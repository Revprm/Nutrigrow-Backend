package seeds

import (
	"errors"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

// ListBahanMakananSeeder menyemai data bahan makanan awal ke dalam database.
func ListBahanMakananSeeder(db *gorm.DB) error {
	bahanMakananList := []entity.BahanMakanan{
		{NamaBahanMakanan: "Telur Ayam", DeskripsiBahan: "Sumber protein hewani yang murah dan mudah didapat."},
		{NamaBahanMakanan: "Bayam", DeskripsiBahan: "Sayuran hijau kaya zat besi dan vitamin."},
		{NamaBahanMakanan: "Daging Sapi", DeskripsiBahan: "Sumber protein dan zat besi tinggi."},
		{NamaBahanMakanan: "Tempe", DeskripsiBahan: "Protein nabati fermentasi yang kaya gizi."},
		{NamaBahanMakanan: "Wortel", DeskripsiBahan: "Kaya vitamin A, baik untuk penglihatan."},
		{NamaBahanMakanan: "Ikan Salmon", DeskripsiBahan: "Sumber Omega-3 dan protein."},
	}

	// Memastikan tabel ada sebelum menyemai data
	hasTable := db.Migrator().HasTable(&entity.BahanMakanan{});
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.BahanMakanan{}); err != nil {
			return err
		}
	}

	for _, data := range bahanMakananList {
		var bahan entity.BahanMakanan
		// Memeriksa apakah data bahan makanan sudah ada berdasarkan nama
		err := db.Where(&entity.BahanMakanan{NamaBahanMakanan: data.NamaBahanMakanan}).First(&bahan).Error
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
