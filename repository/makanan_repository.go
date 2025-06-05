package repository

import (
	"context"
	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

type MakananRepository interface {
	Create(ctx context.Context, tx *gorm.DB, makanan entity.Makanan) (entity.Makanan, error)
	GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.Makanan, error)
	GetByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.Makanan, error)
	GetAllWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) ([]entity.Makanan, int64, error)
	GetByBahanMakanan(ctx context.Context, tx *gorm.DB, bahanID string, req dto.PaginationRequest) ([]entity.Makanan, int64, error)
	Update(ctx context.Context, tx *gorm.DB, makanan entity.Makanan) (entity.Makanan, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
}

type makananRepository struct {
	db *gorm.DB
}

func NewMakananRepository(db *gorm.DB) MakananRepository {
	return &makananRepository{db: db}
}

func (r *makananRepository) Create(ctx context.Context, tx *gorm.DB, makanan entity.Makanan) (entity.Makanan, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&makanan).Error; err != nil {
		return entity.Makanan{}, err
	}
	return makanan, nil
}

func (r *makananRepository) GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.Makanan, error) {
	if tx == nil {
		tx = r.db
	}

	var makanan entity.Makanan
	if err := tx.WithContext(ctx).Preload("BahanMakanans").First(&makanan, "id = ?", id).Error; err != nil {
		return entity.Makanan{}, err
	}
	return makanan, nil
}

func (r *makananRepository) GetByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.Makanan, error) {
	if tx == nil {
		tx = r.db
	}

	var makanan entity.Makanan
	if err := tx.WithContext(ctx).Preload("BahanMakanans").First(&makanan, "nama_makanan = ?", nama).Error; err != nil {
		return entity.Makanan{}, err
	}
	return makanan, nil
}

func (r *makananRepository) GetAllWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) ([]entity.Makanan, int64, error) {
	if tx == nil {
		tx = r.db
	}

	var makanans []entity.Makanan
	var count int64

	query := tx.WithContext(ctx).Preload("BahanMakanans")
	if req.Search != "" {
		query = query.Where("nama_makanan LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Model(&entity.Makanan{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PerPage
	if err := query.Offset(offset).Limit(req.PerPage).Find(&makanans).Error; err != nil {
		return nil, 0, err
	}

	return makanans, count, nil
}

func (r *makananRepository) GetByBahanMakanan(ctx context.Context, tx *gorm.DB, bahanID string, req dto.PaginationRequest) ([]entity.Makanan, int64, error) {
	if tx == nil {
		tx = r.db
	}

	var makanans []entity.Makanan
	var count int64

	query := tx.WithContext(ctx).Preload("BahanMakanans").
		Joins("JOIN makanan_bahan_makanan_pivot ON makanan_bahan_makanan_pivot.makanan_id = makanans.id").
		Where("makanan_bahan_makanan_pivot.bahan_makanan_id = ?", bahanID)

	if req.Search != "" {
		query = query.Where("nama_makanan LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Model(&entity.Makanan{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PerPage
	if err := query.Offset(offset).Limit(req.PerPage).Find(&makanans).Error; err != nil {
		return nil, 0, err
	}

	return makanans, count, nil
}

func (r *makananRepository) Update(ctx context.Context, tx *gorm.DB, makanan entity.Makanan) (entity.Makanan, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&makanan).Error; err != nil {
		return entity.Makanan{}, err
	}
	return makanan, nil
}

func (r *makananRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	// First, load the Makanan item to access its associations
	var makanan entity.Makanan
	if err := tx.WithContext(ctx).Preload("BahanMakanans").First(&makanan, "id = ?", id).Error; err != nil {
		return err // Return error if makanan is not found
	}

	// FIX: Clear the many-to-many association before deleting the main record
	if err := tx.WithContext(ctx).Model(&makanan).Association("BahanMakanans").Clear(); err != nil {
		return err // Return error if clearing association fails
	}

	// Now delete the Makanan record
	if err := tx.WithContext(ctx).Delete(&makanan, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
