package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	// "strconv"
	"bytes"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/entity"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	// "golang.org/x/text/cases"

	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	StuntingService interface {
		Create(ctx context.Context, req dto.StuntingCreateRequest) (dto.StuntingResponse, error)
		GetByID(ctx context.Context, id string) (dto.StuntingResponse, error)
		GetByUserID(ctx context.Context, userID string, req dto.PaginationRequest) (dto.StuntingPaginationResponse, error)
		GetLatestByUserID(ctx context.Context, userID string) (dto.StuntingResponse, error)
		Update(ctx context.Context, id string, req dto.StuntingUpdateRequest) (dto.StuntingResponse, error)
		Delete(ctx context.Context, id string) error
		Predict(ctx context.Context, req dto.StuntingPredictRequest) (dto.StuntingPredictResponse, error)
	}

	stuntingService struct {
		stuntingRepo repository.StuntingRepository
		db           *gorm.DB
		mlApiUrl     string
	}
)

func NewStuntingService(stuntingRepo repository.StuntingRepository, db *gorm.DB, mlApiUrl string) StuntingService {
	return &stuntingService{
		stuntingRepo: stuntingRepo,
		db:           db,
		mlApiUrl:     mlApiUrl,
	}
}

func (s *stuntingService) Create(ctx context.Context, req dto.StuntingCreateRequest) (dto.StuntingResponse, error) {
	// Convert gender from string to int
	jenisKelamin := 0
	if req.JenisKelamin == "Laki-laki" {
		jenisKelamin = 1
	}

	stunting := entity.Stunting{
		UserID:          req.UserID,
		JenisKelamin:    jenisKelamin,
		TinggiBadan:     req.TinggiBadan,
		CatatanStunting: req.CatatanStunting,
	}

	result, err := s.stuntingRepo.Create(ctx, nil, stunting)
	if err != nil {
		return dto.StuntingResponse{}, dto.ErrCreateStunting
	}

	return dto.StuntingResponse{
		ID:              result.ID,
		UserID:          result.UserID,
		JenisKelamin:    req.JenisKelamin,
		TinggiBadan:     result.TinggiBadan,
		CatatanStunting: result.CatatanStunting,
		HasilPrediksi:   result.HasilPrediksi,
	}, nil
}

func (s *stuntingService) GetByID(ctx context.Context, id string) (dto.StuntingResponse, error) {
	result, err := s.stuntingRepo.GetByID(ctx, nil, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.StuntingResponse{}, dto.ErrStuntingNotFound
		}
		return dto.StuntingResponse{}, dto.ErrGetStuntingById
	}

	jenisKelamin := "Perempuan"
	if result.JenisKelamin == 1 {
		jenisKelamin = "Laki-laki"
	}

	return dto.StuntingResponse{
		ID:              result.ID,
		UserID:          result.UserID,
		JenisKelamin:    jenisKelamin,
		TinggiBadan:     result.TinggiBadan,
		CatatanStunting: result.CatatanStunting,
		HasilPrediksi:   result.HasilPrediksi,
	}, nil
}

func (s *stuntingService) GetByUserID(ctx context.Context, userID string, req dto.PaginationRequest) (dto.StuntingPaginationResponse, error) {
	results, count, err := s.stuntingRepo.GetByUserID(ctx, nil, userID, req)
	if err != nil {
		return dto.StuntingPaginationResponse{}, dto.ErrGetStuntingByUserId
	}

	var data []dto.StuntingResponse
	for _, result := range results {
		jenisKelamin := "Perempuan"
		if result.JenisKelamin == 1 {
			jenisKelamin = "Laki-laki"
		}

		data = append(data, dto.StuntingResponse{
			ID:              result.ID,
			UserID:          result.UserID,
			JenisKelamin:    jenisKelamin,
			TinggiBadan:     result.TinggiBadan,
			CatatanStunting: result.CatatanStunting,
			HasilPrediksi:   result.HasilPrediksi,
		})
	}

	return dto.StuntingPaginationResponse{
		Data: data,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
		},
	}, nil
}

func (s *stuntingService) GetLatestByUserID(ctx context.Context, userID string) (dto.StuntingResponse, error) {
	result, err := s.stuntingRepo.GetLatestByUserID(ctx, nil, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.StuntingResponse{}, dto.ErrStuntingNotFound
		}
		return dto.StuntingResponse{}, dto.ErrGetStuntingByUserId
	}

	jenisKelamin := "Perempuan"
	if result.JenisKelamin == 1 {
		jenisKelamin = "Laki-laki"
	}

	return dto.StuntingResponse{
		ID:              result.ID,
		UserID:          result.UserID,
		JenisKelamin:    jenisKelamin,
		TinggiBadan:     result.TinggiBadan,
		CatatanStunting: result.CatatanStunting,
		HasilPrediksi:   result.HasilPrediksi,
	}, nil
}

func (s *stuntingService) Update(ctx context.Context, id string, req dto.StuntingUpdateRequest) (dto.StuntingResponse, error) {
	existing, err := s.stuntingRepo.GetByID(ctx, nil, id)
	if err != nil {
		return dto.StuntingResponse{}, dto.ErrStuntingNotFound
	}

	// Convert gender from string to int if provided
	if req.JenisKelamin != "" {
		if req.JenisKelamin == "Laki-laki" {
			existing.JenisKelamin = 1
		} else {
			existing.JenisKelamin = 0
		}
	}

	if req.TinggiBadan != 0 {
		existing.TinggiBadan = req.TinggiBadan
	}

	if req.CatatanStunting != "" {
		existing.CatatanStunting = req.CatatanStunting
	}

	if req.HasilPrediksi != "" {
		existing.HasilPrediksi = req.HasilPrediksi
	}

	updated, err := s.stuntingRepo.Update(ctx, nil, existing)
	if err != nil {
		return dto.StuntingResponse{}, dto.ErrUpdateStunting
	}

	jenisKelamin := "Perempuan"
	if updated.JenisKelamin == 1 {
		jenisKelamin = "Laki-laki"
	}

	return dto.StuntingResponse{
		ID:              updated.ID,
		UserID:          updated.UserID,
		JenisKelamin:    jenisKelamin,
		TinggiBadan:     updated.TinggiBadan,
		CatatanStunting: updated.CatatanStunting,
		HasilPrediksi:   updated.HasilPrediksi,
	}, nil
}

func (s *stuntingService) Delete(ctx context.Context, id string) error {
	if err := s.stuntingRepo.Delete(ctx, nil, id); err != nil {
		return dto.ErrDeleteStunting
	}
	return nil
}

func (s *stuntingService) Predict(ctx context.Context, req dto.StuntingPredictRequest) (dto.StuntingPredictResponse, error) {
	// Convert gender to int for ML API
	gender := 0
	if req.JenisKelamin == "Laki-laki" {
		gender = 1
	}

	// Prepare request payload for ML API
	payload := map[string]interface{}{
		"Umur (bulan)":         req.UmurBulan,
		"Jenis Kelamin":        gender,
		"Tinggi Badan (cm)":    req.TinggiBadan,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return dto.StuntingPredictResponse{}, fmt.Errorf("failed to marshal prediction payload: %v", err)
	}

	// Make HTTP request to ML API
	resp, err := http.Post(s.mlApiUrl, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return dto.StuntingPredictResponse{}, fmt.Errorf("failed to call prediction API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dto.StuntingPredictResponse{}, fmt.Errorf("prediction API returned status: %d", resp.StatusCode)
	}

	// Read and parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.StuntingPredictResponse{}, fmt.Errorf("failed to read prediction response: %v", err)
	}

	var predictionResponse struct {
		StatusGizi int     `json:"Status_Gizi"`
		Confidence float64 `json:"confidence"`
	}

	if err := json.Unmarshal(body, &predictionResponse); err != nil {
		return dto.StuntingPredictResponse{}, fmt.Errorf("failed to parse prediction response: %v", err)
	}

	// Map status gizi to human-readable format
	var statusGizi string
	switch predictionResponse.StatusGizi {
	case 0:
		statusGizi = "Normal"
	case 1:
		statusGizi = "Stunting Berat"
	case 2:
		statusGizi = "Stunting"
	case 3:
		statusGizi = "Gizi Lebih"
	default:
		statusGizi = "Tidak Diketahui"
	}

	return dto.StuntingPredictResponse{
		StatusGizi: statusGizi,
		Confidence: predictionResponse.Confidence,
	}, nil
}