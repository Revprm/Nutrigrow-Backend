package entity

import (
	"github.com/google/uuid"
)

type KategoriBerita struct {
	ID                 uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	NamaKategoriBerita string    `gorm:"type:varchar(255);not null;unique;column:nama_kategori_berita"`

	Beritas []Berita `gorm:"foreignKey:KategoriBeritaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // A category can have

	Timestamp
}