package seeds

import (
	"errors"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

// ListBeritaSeeder menyemai data berita awal ke dalam database.
func ListBeritaSeeder(db *gorm.DB) error {
	// Ambil kategori berita yang sudah ada untuk relasi
	var kategoriGiziAnak entity.KategoriBerita
	db.Where("nama_kategori_berita = ?", "Gizi Anak").First(&kategoriGiziAnak)

	var kategoriKesehatanIbu entity.KategoriBerita
	db.Where("nama_kategori_berita = ?", "Kesehatan Ibu").First(&kategoriKesehatanIbu)

	var kategoriPencegahanStunting entity.KategoriBerita
	db.Where("nama_kategori_berita = ?", "Pencegahan Stunting").First(&kategoriPencegahanStunting)

	beritaList := []entity.Berita{
		{
			JudulBerita:      "Pentingnya Gizi Seimbang untuk Tumbuh Kembang Anak",
			SubjudulBerita:   "Panduan Lengkap Nutrisi Anak Usia Dini",
			DeskripsiBerita:  "Gizi seimbang adalah kunci utama dalam memastikan tumbuh kembang anak yang optimal. Artikel ini membahas tentang nutrisi esensial yang dibutuhkan anak-anak...",
			KategoriBeritaID: kategoriGiziAnak.ID,
		},
		{
			JudulBerita:      "Peran Asupan Protein dalam Mencegah Stunting",
			SubjudulBerita:   "Sumber Protein Terbaik untuk Balita",
			DeskripsiBerita:  "Stunting masih menjadi masalah serius di banyak negara. Salah satu cara efektif untuk mencegahnya adalah dengan memastikan asupan protein yang cukup...",
			KategoriBeritaID: kategoriPencegahanStunting.ID,
		},
		{
			JudulBerita:      "Tips Menjaga Kesehatan Ibu Hamil dan Menyusui",
			SubjudulBerita:   "Nutrisi dan Gaya Hidup Sehat untuk Ibu",
			DeskripsiBerita:  "Kesehatan ibu adalah fondasi bagi kesehatan anak. Artikel ini memberikan tips praktis untuk menjaga nutrisi dan gaya hidup sehat selama kehamilan dan menyusui...",
			KategoriBeritaID: kategoriKesehatanIbu.ID,
		},
	}

	// Memastikan tabel ada sebelum menyemai data
	hasTable := db.Migrator().HasTable(&entity.Berita{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Berita{}); err != nil {
			return err
		}
	}

	for _, data := range beritaList {
		var berita entity.Berita
		// Memeriksa apakah data berita sudah ada berdasarkan judul
		err := db.Where(&entity.Berita{JudulBerita: data.JudulBerita}).First(&berita).Error
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
