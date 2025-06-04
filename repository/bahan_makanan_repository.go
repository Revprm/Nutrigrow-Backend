package repository

import (
	"context"
	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

type BahanMakananRepository interface {
	Create(ctx context.Context, tx *gorm.DB, bahan entity.BahanMakanan) (entity.BahanMakanan, error)
	GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.BahanMakanan, error)
	GetByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.BahanMakanan, error)
	GetAllWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) ([]entity.BahanMakanan, int64, error)
	Update(ctx context.Context, tx *gorm.DB, bahan entity.BahanMakanan) (entity.BahanMakanan, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
}

type bahanMakananRepository struct {
	db *gorm.DB
}

func NewBahanMakananRepository(db *gorm.DB) BahanMakananRepository {
	return &bahanMakananRepository{db: db}
}

func (r *bahanMakananRepository) Create(ctx context.Context, tx *gorm.DB, bahan entity.BahanMakanan) (entity.BahanMakanan, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&bahan).Error; err != nil {
		return entity.BahanMakanan{}, err
	}
	return bahan, nil
}

func (r *bahanMakananRepository) GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.BahanMakanan, error) {
	if tx == nil {
		tx = r.db
	}

	var bahan entity.BahanMakanan
	if err := tx.WithContext(ctx).Preload("Makanans").First(&bahan, "id = ?", id).Error; err != nil {
		return entity.BahanMakanan{}, err
	}
	return bahan, nil
}

func (r *bahanMakananRepository) GetByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.BahanMakanan, error) {
	if tx == nil {
		tx = r.db
	}

	var bahan entity.BahanMakanan
	if err := tx.WithContext(ctx).Preload("Makanans").First(&bahan, "nama_bahan_makanan = ?", nama).Error; err != nil {
		return entity.BahanMakanan{}, err
	}
	return bahan, nil
}

func (r *bahanMakananRepository) GetAllWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) ([]entity.BahanMakanan, int64, error) {
	if tx == nil {
		tx = r.db
	}

	var bahanList []entity.BahanMakanan
	var count int64

	query := tx.WithContext(ctx).Model(&entity.BahanMakanan{})
	if req.Search != "" {
		query = query.Where("nama_bahan_makanan LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PerPage
	if err := query.Offset(offset).Limit(req.PerPage).Find(&bahanList).Error; err != nil {
		return nil, 0, err
	}

	return bahanList, count, nil
}

func (r *bahanMakananRepository) Update(ctx context.Context, tx *gorm.DB, bahan entity.BahanMakanan) (entity.BahanMakanan, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&bahan).Error; err != nil {
		return entity.BahanMakanan{}, err
	}
	return bahan, nil
}

func (r *bahanMakananRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.BahanMakanan{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}