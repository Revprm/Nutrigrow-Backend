package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	MakananService interface {
		Create(ctx context.Context, req dto.MakananCreateRequest) (dto.MakananResponse, error)
		GetByID(ctx context.Context, id string) (dto.MakananResponse, error)
		GetByNama(ctx context.Context, nama string) (dto.MakananResponse, error)
		GetAllWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MakananPaginationResponse, error)
		GetByBahanMakanan(ctx context.Context, bahanID string, req dto.PaginationRequest) (dto.MakananPaginationResponse, error)
		Update(ctx context.Context, req dto.MakananUpdateRequest, id string) (dto.MakananResponse, error)
		Delete(ctx context.Context, id string) error
	}

	makananService struct {
		makananRepo       repository.MakananRepository
		bahanMakananRepo  repository.BahanMakananRepository
		db                *gorm.DB
	}
)

func NewMakananService(
	makananRepo repository.MakananRepository,
	bahanMakananRepo repository.BahanMakananRepository,
	db *gorm.DB,
) MakananService {
	return &makananService{
		makananRepo:      makananRepo,
		bahanMakananRepo: bahanMakananRepo,
		db:               db,
	}
}

func (s *makananService) Create(ctx context.Context, req dto.MakananCreateRequest) (dto.MakananResponse, error) {
	// Convert UUIDs to strings for repository calls
	bahanMakanans := make([]*entity.BahanMakanan, 0, len(req.BahanMakananIDs))
	for _, id := range req.BahanMakananIDs {
		bahan, err := s.bahanMakananRepo.GetByID(ctx, nil, id.String())
		if err != nil {
			return dto.MakananResponse{}, fmt.Errorf("bahan makanan with ID %s not found: %w", id.String(), err)
		}
		bahanMakanans = append(bahanMakanans, &bahan)
	}

	makanan := entity.Makanan{
		NamaMakanan:          req.NamaMakanan,
		DeskripsiMakanan:     req.DeskripsiMakanan,
		VideoTutorialMakanan: req.VideoTutorialMakanan,
		BahanMakanans:        bahanMakanans,
	}

	result, err := s.makananRepo.Create(ctx, nil, makanan)
	if err != nil {
		return dto.MakananResponse{}, fmt.Errorf("failed to create makanan: %w", err)
	}

	return dto.MakananResponse{
		ID:                   result.ID,
		NamaMakanan:          result.NamaMakanan,
		DeskripsiMakanan:     result.DeskripsiMakanan,
		VideoTutorialMakanan: result.VideoTutorialMakanan,
		BahanMakanans:        convertBahanMakanansToDTO(result.BahanMakanans),
	}, nil
}

func (s *makananService) GetByID(ctx context.Context, id string) (dto.MakananResponse, error) {
	result, err := s.makananRepo.GetByID(ctx, nil, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.MakananResponse{}, dto.ErrNotFound
		}
		return dto.MakananResponse{}, fmt.Errorf("failed to get makanan by ID: %w", err)
	}

	return dto.MakananResponse{
		ID:                   result.ID,
		NamaMakanan:          result.NamaMakanan,
		DeskripsiMakanan:     result.DeskripsiMakanan,
		VideoTutorialMakanan: result.VideoTutorialMakanan,
		BahanMakanans:        convertBahanMakanansToDTO(result.BahanMakanans),
	}, nil
}

func (s *makananService) GetByNama(ctx context.Context, nama string) (dto.MakananResponse, error) {
	result, err := s.makananRepo.GetByNama(ctx, nil, nama)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.MakananResponse{}, dto.ErrNotFound
		}
		return dto.MakananResponse{}, fmt.Errorf("failed to get makanan by name: %w", err)
	}

	return dto.MakananResponse{
		ID:                   result.ID,
		NamaMakanan:          result.NamaMakanan,
		DeskripsiMakanan:     result.DeskripsiMakanan,
		VideoTutorialMakanan: result.VideoTutorialMakanan,
		BahanMakanans:        convertBahanMakanansToDTO(result.BahanMakanans),
	}, nil
}

func (s *makananService) GetAllWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MakananPaginationResponse, error) {
	req.Default() // Apply default pagination values

	result, count, err := s.makananRepo.GetAllWithPagination(ctx, nil, req)
	if err != nil {
		return dto.MakananPaginationResponse{}, fmt.Errorf("failed to get makanan with pagination: %w", err)
	}

	var makananList []dto.MakananResponse
	for _, makanan := range result {
		makananList = append(makananList, dto.MakananResponse{
			ID:                   makanan.ID,
			NamaMakanan:          makanan.NamaMakanan,
			DeskripsiMakanan:     makanan.DeskripsiMakanan,
			VideoTutorialMakanan: makanan.VideoTutorialMakanan,
			BahanMakanans:        convertBahanMakanansToDTO(makanan.BahanMakanans),
		})
	}

	return dto.MakananPaginationResponse{
		Data: makananList,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
		},
	}, nil
}

func (s *makananService) GetByBahanMakanan(ctx context.Context, bahanID string, req dto.PaginationRequest) (dto.MakananPaginationResponse, error) {
	req.Default() // Apply default pagination values

	result, count, err := s.makananRepo.GetByBahanMakanan(ctx, nil, bahanID, req)
	if err != nil {
		return dto.MakananPaginationResponse{}, fmt.Errorf("failed to get makanan by bahan: %w", err)
	}

	var makananList []dto.MakananResponse
	for _, makanan := range result {
		makananList = append(makananList, dto.MakananResponse{
			ID:                   makanan.ID,
			NamaMakanan:          makanan.NamaMakanan,
			DeskripsiMakanan:     makanan.DeskripsiMakanan,
			VideoTutorialMakanan: makanan.VideoTutorialMakanan,
			BahanMakanans:        convertBahanMakanansToDTO(makanan.BahanMakanans),
		})
	}

	return dto.MakananPaginationResponse{
		Data: makananList,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
		},
	}, nil
}

func (s *makananService) Update(ctx context.Context, req dto.MakananUpdateRequest, id string) (dto.MakananResponse, error) {
	makanan, err := s.makananRepo.GetByID(ctx, nil, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.MakananResponse{}, dto.ErrNotFound
		}
		return dto.MakananResponse{}, fmt.Errorf("failed to get makanan for update: %w", err)
	}

	// Update basic fields
	if req.NamaMakanan != "" {
		makanan.NamaMakanan = req.NamaMakanan
	}
	if req.DeskripsiMakanan != "" {
		makanan.DeskripsiMakanan = req.DeskripsiMakanan
	}
	if req.VideoTutorialMakanan != "" {
		makanan.VideoTutorialMakanan = req.VideoTutorialMakanan
	}

	// Update bahan makanan relationships if provided
	if req.BahanMakananIDs != nil {
		bahanMakanans := make([]*entity.BahanMakanan, 0, len(req.BahanMakananIDs))
		for _, bahanID := range req.BahanMakananIDs {
			bahan, err := s.bahanMakananRepo.GetByID(ctx, nil, bahanID.String())
			if err != nil {
				return dto.MakananResponse{}, fmt.Errorf("bahan makanan with ID %s not found: %w", bahanID.String(), err)
			}
			bahanMakanans = append(bahanMakanans, &bahan)
		}
		makanan.BahanMakanans = bahanMakanans
	}

	updated, err := s.makananRepo.Update(ctx, nil, makanan)
	if err != nil {
		return dto.MakananResponse{}, fmt.Errorf("failed to update makanan: %w", err)
	}

	return dto.MakananResponse{
		ID:                   updated.ID,
		NamaMakanan:          updated.NamaMakanan,
		DeskripsiMakanan:     updated.DeskripsiMakanan,
		VideoTutorialMakanan: updated.VideoTutorialMakanan,
		BahanMakanans:        convertBahanMakanansToDTO(updated.BahanMakanans),
	}, nil
}

func (s *makananService) Delete(ctx context.Context, id string) error {
	if err := s.makananRepo.Delete(ctx, nil, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ErrNotFound
		}
		return fmt.Errorf("failed to delete makanan: %w", err)
	}
	return nil
}

func convertBahanMakanansToDTO(bahanMakanans []*entity.BahanMakanan) []dto.BahanMakananSimpleResponse {
	var result []dto.BahanMakananSimpleResponse
	for _, bahan := range bahanMakanans {
		if bahan != nil {
			result = append(result, dto.BahanMakananSimpleResponse{
				ID:               bahan.ID,
				NamaBahanMakanan: bahan.NamaBahanMakanan,
			})
		}
	}
	return result
}