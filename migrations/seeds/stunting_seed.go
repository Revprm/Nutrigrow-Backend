package seeds

import (
	"errors"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
	"time"
)

// ListStuntingSeeder menyemai data catatan stunting awal ke dalam database.
func ListStuntingSeeder(db *gorm.DB) error {
	// Ambil user yang sudah ada untuk relasi
	var user entity.User
	// Asumsi ada user dengan email 'user1234@gmail.com' dari user_seed.go
	err := db.Where("email = ?", "user1234@gmail.com").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika user tidak ditemukan, buat dummy user agar seeder stunting bisa berjalan
			dummyUser := entity.User{
				Name:       "Dummy User",
				Email:      "dummyuser@example.com",
				Password:   "password123", // Password akan di-hash oleh BeforeCreate hook
				Role:       "user",
				IsVerified: true,
			}
			if err := db.Create(&dummyUser).Error; err != nil {
				return err
			}
			user = dummyUser
		} else {
			return err
		}
	}

	stuntingList := []entity.Stunting{
		{
			UserID:          user.ID,
			JenisKelamin:    1, // 1 untuk Laki-laki
			TinggiBadan:     80.5,
			CatatanStunting: "Pengukuran rutin bulan ke-12. Pertumbuhan sesuai standar.",
			HasilPrediksi:   "Normal",
			Timestamp: entity.Timestamp{
				CreatedAt: time.Now().AddDate(0, -2, 0), // 2 bulan lalu
				UpdatedAt: time.Now().AddDate(0, -2, 0),
			},
		},
		{
			UserID:          user.ID,
			JenisKelamin:    0, // 0 untuk Perempuan
			TinggiBadan:     72.0,
			CatatanStunting: "Ada indikasi stunting ringan, perlu perhatian lebih pada asupan gizi.",
			HasilPrediksi:   "Stunting Ringan",
			Timestamp: entity.Timestamp{
				CreatedAt: time.Now().AddDate(0, -1, 0), // 1 bulan lalu
				UpdatedAt: time.Now().AddDate(0, -1, 0),
			},
		},
		{
			UserID:          user.ID,
			JenisKelamin:    1, // 1 untuk Laki-laki
			TinggiBadan:     85.0,
			CatatanStunting: "Perkembangan sangat baik setelah intervensi gizi.",
			HasilPrediksi:   "Normal",
			Timestamp: entity.Timestamp{
				CreatedAt: time.Now(), // Sekarang
				UpdatedAt: time.Now(),
			},
		},
	}

	// Memastikan tabel ada sebelum menyemai data
	hasTable := db.Migrator().HasTable(&entity.Stunting{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Stunting{}); err != nil {
			return err
		}
	}

	for _, data := range stuntingList {
		// Untuk stunting, kita tidak perlu cek duplikasi berdasarkan nama/email karena bisa ada banyak record untuk satu user
		// Cukup buat record baru
		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}

	return nil
}
