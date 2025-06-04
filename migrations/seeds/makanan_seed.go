package seeds

import (
	"errors"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

// ListMakananSeeder menyemai data makanan awal ke dalam database.
func ListMakananSeeder(db *gorm.DB) error {
	// Ambil bahan makanan yang sudah ada untuk relasi
	var telurAyam entity.BahanMakanan
	db.Where("nama_bahan_makanan = ?", "Telur Ayam").First(&telurAyam)

	var bayam entity.BahanMakanan
	db.Where("nama_bahan_makanan = ?", "Bayam").First(&bayam)

	var dagingSapi entity.BahanMakanan
	db.Where("nama_bahan_makanan = ?", "Daging Sapi").First(&dagingSapi)

	var tempe entity.BahanMakanan
	db.Where("nama_bahan_makanan = ?", "Tempe").First(&tempe)

	var wortel entity.BahanMakanan
	db.Where("nama_bahan_makanan = ?", "Wortel").First(&wortel)

	var ikanSalmon entity.BahanMakanan
	db.Where("nama_bahan_makanan = ?", "Ikan Salmon").First(&ikanSalmon)

	makananList := []entity.Makanan{
		{
			NamaMakanan:          "Omelet Sayur Bergizi",
			DeskripsiMakanan:     "Omelet lezat dengan campuran sayuran, cocok untuk sarapan anak.",
			VideoTutorialMakanan: "https://www.youtube.com/watch?v=omelet_sayur",
			BahanMakanans:        []*entity.BahanMakanan{&telurAyam, &bayam},
		},
		{
			NamaMakanan:          "Sup Daging dan Wortel",
			DeskripsiMakanan:     "Sup hangat kaya nutrisi dengan potongan daging sapi dan wortel.",
			VideoTutorialMakanan: "https://www.youtube.com/watch?v=sup_daging",
			BahanMakanans:        []*entity.BahanMakanan{&dagingSapi, &wortel},
		},
		{
			NamaMakanan:          "Nugget Tempe Homemade",
			DeskripsiMakanan:     "Alternatif nugget sehat dari tempe, disukai anak-anak.",
			VideoTutorialMakanan: "https://www.youtube.com/watch?v=nugget_tempe",
			BahanMakanans:        []*entity.BahanMakanan{&tempe},
		},
		{
			NamaMakanan:          "Salmon Panggang Madu",
			DeskripsiMakanan:     "Resep salmon panggang yang lezat dan kaya omega-3.",
			VideoTutorialMakanan: "https://www.youtube.com/watch?v=salmon_panggang",
			BahanMakanans:        []*entity.BahanMakanan{&ikanSalmon},
		},
	}

	// Memastikan tabel ada sebelum menyemai data
	hasTable := db.Migrator().HasTable(&entity.Makanan{});
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Makanan{}); err != nil {
			return err
		}
	}

	for _, data := range makananList {
		var makanan entity.Makanan
		// Memeriksa apakah data makanan sudah ada berdasarkan nama
		err := db.Where(&entity.Makanan{NamaMakanan: data.NamaMakanan}).First(&makanan).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// Jika data belum ada, buat record baru
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		} else {
			// Jika makanan sudah ada, perbarui relasi many-to-many
			if err := db.Model(&makanan).Association("BahanMakanans").Replace(data.BahanMakanans); err != nil {
				return err
			}
		}
	}

	return nil
}
