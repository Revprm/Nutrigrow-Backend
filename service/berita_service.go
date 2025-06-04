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
	BeritaService interface {
		Create(ctx context.Context, req dto.BeritaCreateRequest) (dto.BeritaResponse, error)
		GetByID(ctx context.Context, id string) (dto.BeritaResponse, error)
		GetByKategori(ctx context.Context, kategoriID string, req dto.PaginationRequest) (dto.BeritaPaginationResponse, error)
		GetAllWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BeritaPaginationResponse, error)
		Update(ctx context.Context, req dto.BeritaUpdateRequest, id string) (dto.BeritaResponse, error)
		Delete(ctx context.Context, id string) error
	}

	beritaService struct {
		beritaRepo        repository.BeritaRepository
		kategoriBeritaRepo repository.KategoriBeritaRepository
		db                *gorm.DB
	}
)

func NewBeritaService(
	beritaRepo repository.BeritaRepository,
	kategoriBeritaRepo repository.KategoriBeritaRepository,
	db *gorm.DB,
) BeritaService {
	return &beritaService{
		beritaRepo:        beritaRepo,
		kategoriBeritaRepo: kategoriBeritaRepo,
		db:               db,
	}
}

func (s *beritaService) Create(ctx context.Context, req dto.BeritaCreateRequest) (dto.BeritaResponse, error) {
	// No need to parse UUID since req.KategoriBeritaID is already uuid.UUID
	kategoriID := req.KategoriBeritaID

	// Check if kategori exists - convert UUID to string
	_, err := s.kategoriBeritaRepo.GetByID(ctx, nil, req.KategoriBeritaID.String())
	if err != nil {
		return dto.BeritaResponse{}, errors.New("kategori not found")
	}

	berita := entity.Berita{
		JudulBerita:      req.JudulBerita,
		SubjudulBerita:   req.SubjudulBerita,
		DeskripsiBerita:  req.DeskripsiBerita,
		KategoriBeritaID: kategoriID,
	}

	result, err := s.beritaRepo.Create(ctx, nil, berita)
	if err != nil {
		return dto.BeritaResponse{}, err
	}

	return dto.BeritaResponse{
		ID:               result.ID,
		JudulBerita:      result.JudulBerita,
		SubjudulBerita:   result.SubjudulBerita,
		DeskripsiBerita:  result.DeskripsiBerita,
		KategoriBeritaID: result.KategoriBeritaID,
	}, nil
}

func (s *beritaService) GetByID(ctx context.Context, id string) (dto.BeritaResponse, error) {
	result, err := s.beritaRepo.GetByID(ctx, nil, id)
	if err != nil {
		return dto.BeritaResponse{}, err
	}

	return dto.BeritaResponse{
		ID:               result.ID,
		JudulBerita:      result.JudulBerita,
		SubjudulBerita:   result.SubjudulBerita,
		DeskripsiBerita:  result.DeskripsiBerita,
		KategoriBeritaID: result.KategoriBeritaID,
		KategoriBerita: dto.KategoriBeritaResponse{
			ID:                 result.KategoriBerita.ID,
			NamaKategoriBerita: result.KategoriBerita.NamaKategoriBerita,
		},
	}, nil
}

func (s *beritaService) GetByKategori(ctx context.Context, kategoriID string, req dto.PaginationRequest) (dto.BeritaPaginationResponse, error) {
	result, count, err := s.beritaRepo.GetByKategori(ctx, nil, kategoriID, req)
	if err != nil {
		return dto.BeritaPaginationResponse{}, err
	}

	var beritaList []dto.BeritaResponse
	for _, berita := range result {
		beritaList = append(beritaList, dto.BeritaResponse{
			ID:               berita.ID,
			JudulBerita:      berita.JudulBerita,
			SubjudulBerita:   berita.SubjudulBerita,
			DeskripsiBerita:  berita.DeskripsiBerita,
			KategoriBeritaID: berita.KategoriBeritaID,
			KategoriBerita: dto.KategoriBeritaResponse{
				ID:                 berita.KategoriBerita.ID,
				NamaKategoriBerita: berita.KategoriBerita.NamaKategoriBerita,
			},
		})
	}

	return dto.BeritaPaginationResponse{
		Data: beritaList,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
		},
	}, nil
}

func (s *beritaService) GetAllWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BeritaPaginationResponse, error) {
	result, count, err := s.beritaRepo.GetAllWithPagination(ctx, nil, req)
	if err != nil {
		return dto.BeritaPaginationResponse{}, err
	}

	var beritaList []dto.BeritaResponse
	for _, berita := range result {
		beritaList = append(beritaList, dto.BeritaResponse{
			ID:               berita.ID,
			JudulBerita:      berita.JudulBerita,
			SubjudulBerita:   berita.SubjudulBerita,
			DeskripsiBerita:  berita.DeskripsiBerita,
			KategoriBeritaID: berita.KategoriBeritaID,
			KategoriBerita: dto.KategoriBeritaResponse{
				ID:                 berita.KategoriBerita.ID,
				NamaKategoriBerita: berita.KategoriBerita.NamaKategoriBerita,
			},
		})
	}

	return dto.BeritaPaginationResponse{
		Data: beritaList,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
		},
	}, nil
}

func (s *beritaService) Update(ctx context.Context, req dto.BeritaUpdateRequest, id string) (dto.BeritaResponse, error) {
	berita, err := s.beritaRepo.GetByID(ctx, nil, id)
	if err != nil {
		return dto.BeritaResponse{}, err
	}

	// Update fields
	berita.JudulBerita = req.JudulBerita
	berita.SubjudulBerita = req.SubjudulBerita
	berita.DeskripsiBerita = req.DeskripsiBerita

	// Update kategori if provided (check if UUID is not zero value)
	if req.KategoriBeritaID != uuid.Nil {
		// Check if new kategori exists
		_, err = s.kategoriBeritaRepo.GetByID(ctx, nil, req.KategoriBeritaID.String())
		if err != nil {
			return dto.BeritaResponse{}, errors.New("kategori not found")
		}

		berita.KategoriBeritaID = req.KategoriBeritaID
	}

	updated, err := s.beritaRepo.Update(ctx, nil, berita)
	if err != nil {
		return dto.BeritaResponse{}, err
	}

	return dto.BeritaResponse{
		ID:               updated.ID,
		JudulBerita:      updated.JudulBerita,
		SubjudulBerita:   updated.SubjudulBerita,
		DeskripsiBerita:  updated.DeskripsiBerita,
		KategoriBeritaID: updated.KategoriBeritaID,
	}, nil
}

func (s *beritaService) Delete(ctx context.Context, id string) error {
	if err := s.beritaRepo.Delete(ctx, nil, id); err != nil {
		return err
	}
	return nil
}