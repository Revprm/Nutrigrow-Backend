package repository

import (
	"context"
	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

type StuntingRepository interface {
	Create(ctx context.Context, tx *gorm.DB, stunting entity.Stunting) (entity.Stunting, error)
	GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.Stunting, error)
	GetByUserID(ctx context.Context, tx *gorm.DB, userID string, req dto.PaginationRequest) ([]entity.Stunting, int64, error)
	GetLatestByUserID(ctx context.Context, tx *gorm.DB, userID string) (entity.Stunting, error)
	Update(ctx context.Context, tx *gorm.DB, stunting entity.Stunting) (entity.Stunting, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
}

type stuntingRepository struct {
	db *gorm.DB
}

func NewStuntingRepository(db *gorm.DB) StuntingRepository {
	return &stuntingRepository{db: db}
}

func (r *stuntingRepository) Create(ctx context.Context, tx *gorm.DB, stunting entity.Stunting) (entity.Stunting, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&stunting).Error; err != nil {
		return entity.Stunting{}, err
	}
	return stunting, nil
}

func (r *stuntingRepository) GetByID(ctx context.Context, tx *gorm.DB, id string) (entity.Stunting, error) {
	if tx == nil {
		tx = r.db
	}

	var stunting entity.Stunting
	if err := tx.WithContext(ctx).Preload("User").First(&stunting, "id = ?", id).Error; err != nil {
		return entity.Stunting{}, err
	}
	return stunting, nil
}

func (r *stuntingRepository) GetByUserID(ctx context.Context, tx *gorm.DB, userID string, req dto.PaginationRequest) ([]entity.Stunting, int64, error) {
	if tx == nil {
		tx = r.db
	}

	var stuntingList []entity.Stunting
	var count int64

	query := tx.WithContext(ctx).Preload("User").Where("user_id = ?", userID)
	if req.Search != "" {
		query = query.Where("hasil_prediksi LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Model(&entity.Stunting{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PerPage
	if err := query.Order("created_at DESC").Offset(offset).Limit(req.PerPage).Find(&stuntingList).Error; err != nil {
		return nil, 0, err
	}

	return stuntingList, count, nil
}

func (r *stuntingRepository) GetLatestByUserID(ctx context.Context, tx *gorm.DB, userID string) (entity.Stunting, error) {
	if tx == nil {
		tx = r.db
	}

	var stunting entity.Stunting
	if err := tx.WithContext(ctx).Preload("User").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		First(&stunting).Error; err != nil {
		return entity.Stunting{}, err
	}
	return stunting, nil
}

func (r *stuntingRepository) Update(ctx context.Context, tx *gorm.DB, stunting entity.Stunting) (entity.Stunting, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&stunting).Error; err != nil {
		return entity.Stunting{}, err
	}
	return stunting, nil
}

func (r *stuntingRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Stunting{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}