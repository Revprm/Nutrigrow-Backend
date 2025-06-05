package service

import (
	"context"
	// "errors"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	"gorm.io/gorm"
)

type (
	BahanMakananService interface {
		Create(ctx context.Context, req dto.BahanMakananCreateRequest) (dto.BahanMakananResponse, error)
		GetByID(ctx context.Context, id string) (dto.BahanMakananResponse, error)
		GetByNama(ctx context.Context, nama string) (dto.BahanMakananResponse, error)
		GetAllWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BahanMakananPaginationResponse, error)
		Update(ctx context.Context, req dto.BahanMakananUpdateRequest, id string) (dto.BahanMakananResponse, error)
		Delete(ctx context.Context, id string) error
	}

	bahanMakananService struct {
		bahanMakananRepo repository.BahanMakananRepository
		db               *gorm.DB
	}
)

func NewBahanMakananService(bahanMakananRepo repository.BahanMakananRepository, db *gorm.DB) BahanMakananService {
	return &bahanMakananService{
		bahanMakananRepo: bahanMakananRepo,
		db:               db,
	}
}

func (s *bahanMakananService) Create(ctx context.Context, req dto.BahanMakananCreateRequest) (dto.BahanMakananResponse, error) {
	bahan := entity.BahanMakanan{
		NamaBahanMakanan: req.NamaBahanMakanan,
		DeskripsiBahan:   req.DeskripsiBahan,
	}

	result, err := s.bahanMakananRepo.Create(ctx, nil, bahan)
	if err != nil {
		return dto.BahanMakananResponse{}, err
	}

	return dto.BahanMakananResponse{
		ID:               result.ID,
		NamaBahanMakanan: result.NamaBahanMakanan,
		DeskripsiBahan:   result.DeskripsiBahan,
	}, nil
}

func (s *bahanMakananService) GetByID(ctx context.Context, id string) (dto.BahanMakananResponse, error) {
	result, err := s.bahanMakananRepo.GetByID(ctx, nil, id)
	if err != nil {
		return dto.BahanMakananResponse{}, err
	}

	return dto.BahanMakananResponse{
		ID:               result.ID,
		NamaBahanMakanan: result.NamaBahanMakanan,
		DeskripsiBahan:   result.DeskripsiBahan,
	}, nil
}

func (s *bahanMakananService) GetByNama(ctx context.Context, nama string) (dto.BahanMakananResponse, error) {
	result, err := s.bahanMakananRepo.GetByNama(ctx, nil, nama)
	if err != nil {
		return dto.BahanMakananResponse{}, err
	}

	return dto.BahanMakananResponse{
		ID:               result.ID,
		NamaBahanMakanan: result.NamaBahanMakanan,
		DeskripsiBahan:   result.DeskripsiBahan,
	}, nil
}

func (s *bahanMakananService) GetAllWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BahanMakananPaginationResponse, error) {
	req.Default() // Apply default pagination values (Page=1, PerPage=10 if not set)

	result, count, err := s.bahanMakananRepo.GetAllWithPagination(ctx, nil, req)
	if err != nil {
		return dto.BahanMakananPaginationResponse{}, err
	}

	var bahanList []dto.BahanMakananResponse
	for _, bahan := range result {
		bahanList = append(bahanList, dto.BahanMakananResponse{
			ID:               bahan.ID,
			NamaBahanMakanan: bahan.NamaBahanMakanan,
			DeskripsiBahan:   bahan.DeskripsiBahan,
		})
	}

	maxPage := repository.TotalPage(count, int64(req.PerPage))

	return dto.BahanMakananPaginationResponse{
		Data: bahanList,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: maxPage,
		},
	}, nil
}

func (s *bahanMakananService) Update(ctx context.Context, req dto.BahanMakananUpdateRequest, id string) (dto.BahanMakananResponse, error) {
	bahan, err := s.bahanMakananRepo.GetByID(ctx, nil, id)
	if err != nil {
		return dto.BahanMakananResponse{}, err
	}

	bahan.NamaBahanMakanan = req.NamaBahanMakanan
	bahan.DeskripsiBahan = req.DeskripsiBahan

	updated, err := s.bahanMakananRepo.Update(ctx, nil, bahan)
	if err != nil {
		return dto.BahanMakananResponse{}, err
	}

	return dto.BahanMakananResponse{
		ID:               updated.ID,
		NamaBahanMakanan: updated.NamaBahanMakanan,
		DeskripsiBahan:   updated.DeskripsiBahan,
	}, nil
}

func (s *bahanMakananService) Delete(ctx context.Context, id string) error {
	if err := s.bahanMakananRepo.Delete(ctx, nil, id); err != nil {
		return err
	}
	return nil
}