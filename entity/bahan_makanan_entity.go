package entity

import (
	"github.com/google/uuid"
)

type BahanMakanan struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	NamaBahanMakanan  string    `gorm:"type:varchar(255);not null;unique;column:nama_bahan_makanan"`
	DeskripsiBahan    string    `gorm:"type:text;column:deskripsi_bahan"`

	Makanans []*Makanan `gorm:"many2many:makanan_bahan_makanan_pivot;foreignKey:ID;joinForeignKey:BahanMakananID;References:ID;joinReferences:MakananID"`

	Timestamp
}