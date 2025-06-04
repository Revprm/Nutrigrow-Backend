package entity

import (
	"github.com/google/uuid"
)

type Berita struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	JudulBerita      string    `gorm:"type:varchar(255);not null;column:judul_berita"`
	SubjudulBerita   string    `gorm:"type:varchar(255);column:subjudul_berita"`
	DeskripsiBerita  string    `gorm:"type:text;not null;column:deskripsi_berita"`
	KategoriBeritaID uuid.UUID `gorm:"type:uuid;not null;index"` // Foreign Key to KategoriBerita

	KategoriBerita KategoriBerita `gorm:"foreignKey:KategoriBeritaID;references:ID"` 

	Timestamp
}