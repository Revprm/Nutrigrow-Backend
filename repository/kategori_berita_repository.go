package repository

import (
	"context"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

type KategoriBeritaRepository interface {
	Create(ctx context.Context, tx *gorm.DB, kategori entity.KategoriBerita) (entity.KategoriBerita, error)
	GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.KategoriBerita, error)
	GetByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.KategoriBerita, error)
	GetAll(ctx context.Context, tx *gorm.DB) ([]entity.KategoriBerita, error)
	Update(ctx context.Context, tx *gorm.DB, kategori entity.KategoriBerita) (entity.KategoriBerita, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
}

type kategoriBeritaRepository struct {
	db *gorm.DB
}

func NewKategoriBeritaRepository(db *gorm.DB) KategoriBeritaRepository {
	return &kategoriBeritaRepository{db: db}
}

func (r *kategoriBeritaRepository) Create(ctx context.Context, tx *gorm.DB, kategori entity.KategoriBerita) (entity.KategoriBerita, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&kategori).Error; err != nil {
		return entity.KategoriBerita{}, err
	}
	return kategori, nil
}

func (r *kategoriBeritaRepository) GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.KategoriBerita, error) {
	if tx == nil {
		tx = r.db
	}

	var kategori entity.KategoriBerita
	if err := tx.WithContext(ctx).Preload("Beritas").First(&kategori, "id = ?", id).Error; err != nil {
		return entity.KategoriBerita{}, err
	}
	return kategori, nil
}

func (r *kategoriBeritaRepository) GetByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.KategoriBerita, error) {
	if tx == nil {
		tx = r.db
	}

	var kategori entity.KategoriBerita
	if err := tx.WithContext(ctx).Preload("Beritas").First(&kategori, "nama_kategori_berita = ?", nama).Error; err != nil {
		return entity.KategoriBerita{}, err
	}
	return kategori, nil
}

func (r *kategoriBeritaRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]entity.KategoriBerita, error) {
	if tx == nil {
		tx = r.db
	}

	var kategories []entity.KategoriBerita
	if err := tx.WithContext(ctx).Find(&kategories).Error; err != nil {
		return nil, err
	}
	return kategories, nil
}

func (r *kategoriBeritaRepository) Update(ctx context.Context, tx *gorm.DB, kategori entity.KategoriBerita) (entity.KategoriBerita, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&kategori).Error; err != nil {
		return entity.KategoriBerita{}, err
	}
	return kategori, nil
}

func (r *kategoriBeritaRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.KategoriBerita{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}