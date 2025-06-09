package entity

import (
	"github.com/Caknoooo/go-gin-clean-starter/helpers" // Assuming this path is correct from your imports
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string    `gorm:"type:varchar(100);not null"`
	Email      string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	TelpNumber string    `gorm:"type:varchar(20);index"` // PDM specified 13, yours has 20. Keeping 20.
	Password   string    `gorm:"type:varchar(255);not null"`
	Role       string    `gorm:"type:varchar(50);not null;default:'user'"`
	ImageUrl   string    `gorm:"type:varchar(255)"`
	IsVerified bool      `gorm:"default:false"`
	Alamat     string    `gorm:"type:varchar(255);column:alamat"` // New from PDM

	// Relationships
	Stuntings []Stunting `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // User can have multiple stunting records

	Timestamp
}

// BeforeCreate hook to hash password and set defaults
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	// Hash password
	if u.Password != "" {
		u.Password, err = helpers.HashPassword(u.Password)
		if err != nil {
			return err
		}
	}

	// Ensure UUID is set
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	// Set default role if not specified
	if u.Role == "" {
		u.Role = "user"
	}

	return nil
}

// BeforeUpdate hook to handle password updates
func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	// Only hash password if it has been changed
	if u.Password != "" {
		u.Password, err = helpers.HashPassword(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
