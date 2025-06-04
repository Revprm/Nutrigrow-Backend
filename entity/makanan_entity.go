package entity

import (
	"github.com/google/uuid"
)

type Makanan struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	NamaMakanan          string    `gorm:"type:varchar(255);not null;column:nama_makanan"`
	DeskripsiMakanan     string    `gorm:"type:text;column:deskripsi_makanan"`
	VideoTutorialMakanan string    `gorm:"type:varchar(255);column:video_tutorial_makanan"` // URL to video

	BahanMakanans []*BahanMakanan `gorm:"many2many:makanan_bahan_makanan_pivot;foreignKey:ID;joinForeignKey:MakananID;References:ID;joinReferences:BahanMakananID"`

	Timestamp
}