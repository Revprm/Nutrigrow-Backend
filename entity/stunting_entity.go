package entity

import (
	"github.com/google/uuid"
)

type Stunting struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID          uuid.UUID `gorm:"type:uuid;not null;index"` 
	JenisKelamin    int       `gorm:"type:int;column:jenis_kelamin"`
	TinggiBadan     float64   `gorm:"type:numeric(5,2);column:tinggi_badan"`    
	CatatanStunting string    `gorm:"type:text;column:catatan_stunting"`
	HasilPrediksi   string    `gorm:"type:varchar(50);column:hasil_prediksi"`

	User User `gorm:"foreignKey:UserID;references:ID"` 

	Timestamp
}