package repository

import (
	"context"
	"time"

	// Corrected import path for the entity package
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"gorm.io/gorm"
)

// RefreshTokenRepository defines the interface for refresh token database operations.
type RefreshTokenRepository interface {
	Create(ctx context.Context, tx *gorm.DB, token entity.RefreshToken) (entity.RefreshToken, error)
	FindByToken(ctx context.Context, tx *gorm.DB, token string) (entity.RefreshToken, error)
	DeleteByUserID(ctx context.Context, tx *gorm.DB, userID string) error
	DeleteByToken(ctx context.Context, tx *gorm.DB, token string) error
	DeleteExpired(ctx context.Context, tx *gorm.DB) error
}

type refreshTokenRepository struct {
	db *gorm.DB
}

// NewRefreshTokenRepository creates a new instance of RefreshTokenRepository.
func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{
		db: db,
	}
}

// Create stores a new refresh token in the database.
func (r *refreshTokenRepository) Create(
	ctx context.Context,
	tx *gorm.DB,
	token entity.RefreshToken,
) (entity.RefreshToken, error) {
	// Use the provided transaction tx if not nil, otherwise use the repository's db.
	db := r.db
	if tx != nil {
		db = tx
	}

	// Create the refresh token record in the database.
	if err := db.WithContext(ctx).Create(&token).Error; err != nil {
		return entity.RefreshToken{}, err
	}

	return token, nil
}

// FindByToken retrieves a refresh token from the database by its token string.
// It also preloads the associated User.
func (r *refreshTokenRepository) FindByToken(ctx context.Context, tx *gorm.DB, token string) (
	entity.RefreshToken,
	error,
) {
	// Use the provided transaction tx if not nil, otherwise use the repository's db.
	db := r.db
	if tx != nil {
		db = tx
	}

	var refreshToken entity.RefreshToken
	// Find the refresh token by the token string and preload the User.
	// .Take returns an error if no record is found.
	if err := db.WithContext(ctx).Where("token = ?", token).Preload("User").Take(&refreshToken).Error; err != nil {
		return entity.RefreshToken{}, err
	}

	return refreshToken, nil
}

// DeleteByUserID removes all refresh tokens associated with a specific user ID.
func (r *refreshTokenRepository) DeleteByUserID(ctx context.Context, tx *gorm.DB, userID string) error {
	// Use the provided transaction tx if not nil, otherwise use the repository's db.
	db := r.db
	if tx != nil {
		db = tx
	}

	// Delete refresh tokens where user_id matches.
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Delete(&entity.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteByToken removes a specific refresh token from the database by its token string.
func (r *refreshTokenRepository) DeleteByToken(ctx context.Context, tx *gorm.DB, token string) error {
	// Use the provided transaction tx if not nil, otherwise use the repository's db.
	db := r.db
	if tx != nil {
		db = tx
	}

	// Delete the refresh token where the token string matches.
	if err := db.WithContext(ctx).Where("token = ?", token).Delete(&entity.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteExpired removes all refresh tokens that have passed their expiration time.
func (r *refreshTokenRepository) DeleteExpired(ctx context.Context, tx *gorm.DB) error {
	// Use the provided transaction tx if not nil, otherwise use the repository's db.
	db := r.db
	if tx != nil {
		db = tx
	}

	// Delete refresh tokens where expires_at is before the current time.
	if err := db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&entity.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}
