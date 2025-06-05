package dto

import (
	"errors"
	"github.com/google/uuid"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_STUNTING      = "failed create stunting record"
	MESSAGE_FAILED_GET_STUNTING         = "failed get stunting record"
	MESSAGE_FAILED_GET_LIST_STUNTING    = "failed get list stunting record"
	MESSAGE_FAILED_UPDATE_STUNTING      = "failed update stunting record"
	MESSAGE_FAILED_DELETE_STUNTING      = "failed delete stunting record"
	MESSAGE_FAILED_PREDICT_STUNTING     = "failed predict stunting"
	MESSAGE_FAILED_GET_STUNTING_BY_USER = "failed get stunting records by user"

	// Success
	MESSAGE_SUCCESS_CREATE_STUNTING      = "success create stunting record"
	MESSAGE_SUCCESS_GET_STUNTING         = "success get stunting record"
	MESSAGE_SUCCESS_GET_LIST_STUNTING    = "success get list stunting record"
	MESSAGE_SUCCESS_UPDATE_STUNTING      = "success update stunting record"
	MESSAGE_SUCCESS_DELETE_STUNTING      = "success delete stunting record"
	MESSAGE_SUCCESS_PREDICT_STUNTING     = "success predict stunting"
)

var (
	ErrCreateStunting             = errors.New("failed to create stunting record")
	ErrGetStuntingById            = errors.New("failed to get stunting by id")
	ErrGetStuntingByUserId        = errors.New("failed to get stunting by user id")
	ErrStuntingNotFound           = errors.New("stunting record not found")
	ErrUpdateStunting             = errors.New("failed to update stunting record")
	ErrDeleteStunting             = errors.New("failed to delete stunting record")
	ErrStuntingInvalidData        = errors.New("invalid stunting data")
	ErrStuntingAlreadyExists      = errors.New("stunting record already exists")
)

type (
	StuntingCreateRequest struct {
		UserID          uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
		JenisKelamin    string    `json:"jenis_kelamin" form:"jenis_kelamin" binding:"required"`
		UmurBulan       int       `json:"umur_bulan" form:"umur_bulan" binding:"required"`
		TinggiBadan     float64   `json:"tinggi_badan" form:"tinggi_badan" binding:"required"` // Keep float64 for database storage if needed
		CatatanStunting string    `json:"catatan_stunting" form:"catatan_stunting"`
	}

	StuntingUpdateRequest struct {
		JenisKelamin    string  `json:"jenis_kelamin" form:"jenis_kelamin"`
		TinggiBadan     float64 `json:"tinggi_badan" form:"tinggi_badan"` // Keep float64 for database storage if needed
		CatatanStunting string  `json:"catatan_stunting" form:"catatan_stunting"`
		HasilPrediksi   string  `json:"hasil_prediksi" form:"hasil_prediksi"`
	}

	StuntingResponse struct {
		ID              uuid.UUID     `json:"id"`
		UserID          uuid.UUID     `json:"user_id"`
		JenisKelamin    string        `json:"jenis_kelamin"`
		TinggiBadan     float64       `json:"tinggi_badan"`
		CatatanStunting string        `json:"catatan_stunting"`
		HasilPrediksi   string        `json:"hasil_prediksi"`
		User            UserResponse  `json:"user,omitempty"`
	}
	StuntingPredictRequest struct {
		UmurBulan    int     `json:"umur_bulan" form:"umur_bulan" binding:"required"`
		JenisKelamin string  `json:"jenis_kelamin" form:"jenis_kelamin" binding:"required"`
		TinggiBadan  int     `json:"tinggi_badan" form:"tinggi_badan" binding:"required"` 
	}

	StuntingPredictResponse struct {
		StatusGizi string  `json:"status_gizi"`
		Confidence float64 `json:"confidence"`
	}

	StuntingPaginationResponse struct {
		Data []StuntingResponse `json:"data"`
		PaginationResponse
	}
)
