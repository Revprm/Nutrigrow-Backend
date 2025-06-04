package service

import (
	"context"
	"errors"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	KategoriBeritaService interface {
		Create(ctx context.Context, req dto.KategoriBeritaCreateRequest) (dto.KategoriBeritaResponse, error)
		GetByID(ctx context.Context, id string) (dto.KategoriBeritaResponse, error)
		GetByNama(ctx context.Context, nama string) (dto.KategoriBeritaResponse, error)
		GetAll(ctx context.Context) ([]dto.KategoriBeritaResponse, error)
		Update(ctx context.Context, req dto.KategoriBeritaUpdateRequest, id string) (dto.KategoriBeritaResponse, error)
		Delete(ctx context.Context, id string) error
	}

	kategoriBeritaService struct {
		kategoriBeritaRepo repository.KategoriBeritaRepository
		db                 *gorm.DB
	}
)

func NewKategoriBeritaService(
	kategoriBeritaRepo repository.KategoriBeritaRepository,
	db *gorm.DB,
) KategoriBeritaService {
	return &kategoriBeritaService{
		kategoriBeritaRepo: kategoriBeritaRepo,
		db:                 db,
	}
}

func (s *kategoriBeritaService) Create(ctx context.Context, req dto.KategoriBeritaCreateRequest) (dto.KategoriBeritaResponse, error) {
	// Check if kategori with same name already exists
	existingKategori, err := s.kategoriBeritaRepo.GetByNama(ctx, nil, req.NamaKategoriBerita)
	if err == nil && existingKategori.ID != uuid.Nil {
		return dto.KategoriBeritaResponse{}, errors.New("kategori with this name already exists")
	}

	kategori := entity.KategoriBerita{
		NamaKategoriBerita: req.NamaKategoriBerita,
	}

	result, err := s.kategoriBeritaRepo.Create(ctx, nil, kategori)
	if err != nil {
		return dto.KategoriBeritaResponse{}, err
	}

	return dto.KategoriBeritaResponse{
		ID:                 result.ID,
		NamaKategoriBerita: result.NamaKategoriBerita,
	}, nil
}

func (s *kategoriBeritaService) GetByID(ctx context.Context, id string) (dto.KategoriBeritaResponse, error) {
	result, err := s.kategoriBeritaRepo.GetByID(ctx, nil, id)
	if err != nil {
		return dto.KategoriBeritaResponse{}, err
	}

	return dto.KategoriBeritaResponse{
		ID:                 result.ID,
		NamaKategoriBerita: result.NamaKategoriBerita,
	}, nil
}

func (s *kategoriBeritaService) GetByNama(ctx context.Context, nama string) (dto.KategoriBeritaResponse, error) {
	result, err := s.kategoriBeritaRepo.GetByNama(ctx, nil, nama)
	if err != nil {
		return dto.KategoriBeritaResponse{}, err
	}

	return dto.KategoriBeritaResponse{
		ID:                 result.ID,
		NamaKategoriBerita: result.NamaKategoriBerita,
	}, nil
}

func (s *kategoriBeritaService) GetAll(ctx context.Context) ([]dto.KategoriBeritaResponse, error) {
	result, err := s.kategoriBeritaRepo.GetAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	var kategoriList []dto.KategoriBeritaResponse
	for _, kategori := range result {
		kategoriList = append(kategoriList, dto.KategoriBeritaResponse{
			ID:                 kategori.ID,
			NamaKategoriBerita: kategori.NamaKategoriBerita,
		})
	}

	return kategoriList, nil
}

func (s *kategoriBeritaService) Update(ctx context.Context, req dto.KategoriBeritaUpdateRequest, id string) (dto.KategoriBeritaResponse, error) {
	kategori, err := s.kategoriBeritaRepo.GetByID(ctx, nil, id)
	if err != nil {
		return dto.KategoriBeritaResponse{}, err
	}

	// Check if new name already exists (excluding current record)
	if req.NamaKategoriBerita != "" && req.NamaKategoriBerita != kategori.NamaKategoriBerita {
		existingKategori, err := s.kategoriBeritaRepo.GetByNama(ctx, nil, req.NamaKategoriBerita)
		if err == nil && existingKategori.ID != kategori.ID {
			return dto.KategoriBeritaResponse{}, errors.New("kategori with this name already exists")
		}
	}

	// Update fields
	if req.NamaKategoriBerita != "" {
		kategori.NamaKategoriBerita = req.NamaKategoriBerita
	}

	updated, err := s.kategoriBeritaRepo.Update(ctx, nil, kategori)
	if err != nil {
		return dto.KategoriBeritaResponse{}, err
	}

	return dto.KategoriBeritaResponse{
		ID:                 updated.ID,
		NamaKategoriBerita: updated.NamaKategoriBerita,
	}, nil
}

func (s *kategoriBeritaService) Delete(ctx context.Context, id string) error {
	// Check if kategori exists
	_, err := s.kategoriBeritaRepo.GetByID(ctx, nil, id)
	if err != nil {
		return errors.New("kategori not found")
	}

	if err := s.kategoriBeritaRepo.Delete(ctx, nil, id); err != nil {
		return err
	}
	return nil
}
