package repository

import (
	"context"
	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

type BeritaRepository interface {
	Create(ctx context.Context, tx *gorm.DB, berita entity.Berita) (entity.Berita, error)
	GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.Berita, error)
	GetByKategori(ctx context.Context, tx *gorm.DB, kategoriID string, req dto.PaginationRequest) ([]entity.Berita, int64, error)
	GetAllWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) ([]entity.Berita, int64, error)
	Update(ctx context.Context, tx *gorm.DB, berita entity.Berita) (entity.Berita, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
}

type beritaRepository struct {
	db *gorm.DB
}

func NewBeritaRepository(db *gorm.DB) BeritaRepository {
	return &beritaRepository{db: db}
}

func (r *beritaRepository) Create(ctx context.Context, tx *gorm.DB, berita entity.Berita) (entity.Berita, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&berita).Error; err != nil {
		return entity.Berita{}, err
	}
	return berita, nil
}

func (r *beritaRepository) GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.Berita, error) {
	if tx == nil {
		tx = r.db
	}

	var berita entity.Berita
	if err := tx.WithContext(ctx).Preload("KategoriBerita").First(&berita, "id = ?", id).Error; err != nil {
		return entity.Berita{}, err
	}
	return berita, nil
}

func (r *beritaRepository) GetByKategori(ctx context.Context, tx *gorm.DB, kategoriID string, req dto.PaginationRequest) ([]entity.Berita, int64, error) {
	if tx == nil {
		tx = r.db
	}

	var beritas []entity.Berita
	var count int64

	query := tx.WithContext(ctx).Preload("KategoriBerita").Where("kategori_berita_id = ?", kategoriID)
	if req.Search != "" {
		query = query.Where("judul_berita LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Model(&entity.Berita{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PerPage
	if err := query.Offset(offset).Limit(req.PerPage).Find(&beritas).Error; err != nil {
		return nil, 0, err
	}

	return beritas, count, nil
}

func (r *beritaRepository) GetAllWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) ([]entity.Berita, int64, error) {
	if tx == nil {
		tx = r.db
	}

	var beritas []entity.Berita
	var count int64

	query := tx.WithContext(ctx).Preload("KategoriBerita")
	if req.Search != "" {
		query = query.Where("judul_berita LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Model(&entity.Berita{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PerPage
	if err := query.Offset(offset).Limit(req.PerPage).Find(&beritas).Error; err != nil {
		return nil, 0, err
	}

	return beritas, count, nil
}

func (r *beritaRepository) Update(ctx context.Context, tx *gorm.DB, berita entity.Berita) (entity.Berita, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&berita).Error; err != nil {
		return entity.Berita{}, err
	}
	return berita, nil
}

func (r *beritaRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Berita{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}