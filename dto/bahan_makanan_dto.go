package dto

import (
	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_CREATE_BAHAN_MAKANAN   = "failed create bahan makanan"
	MESSAGE_FAILED_GET_BAHAN_MAKANAN      = "failed get bahan makanan"
	MESSAGE_FAILED_UPDATE_BAHAN_MAKANAN   = "failed update bahan makanan"
	MESSAGE_FAILED_DELETE_BAHAN_MAKANAN   = "failed delete bahan makanan"
	MESSAGE_FAILED_LIST_BAHAN_MAKANAN     = "failed get list bahan makanan"
	MESSAGE_SUCCESS_CREATE_BAHAN_MAKANAN  = "success create bahan makanan"
	MESSAGE_SUCCESS_GET_BAHAN_MAKANAN     = "success get bahan makanan"
	MESSAGE_SUCCESS_UPDATE_BAHAN_MAKANAN  = "success update bahan makanan"
	MESSAGE_SUCCESS_DELETE_BAHAN_MAKANAN  = "success delete bahan makanan"
	MESSAGE_SUCCESS_LIST_BAHAN_MAKANAN    = "success get list bahan makanan"
)

type BahanMakananCreateRequest struct {
	NamaBahanMakanan string `json:"nama_bahan_makanan" binding:"required"`
	DeskripsiBahan   string `json:"deskripsi_bahan"`
}

type BahanMakananUpdateRequest struct {
	NamaBahanMakanan string `json:"nama_bahan_makanan"`
	DeskripsiBahan   string `json:"deskripsi_bahan"`
}

type BahanMakananResponse struct {
	ID               uuid.UUID `json:"id"`
	NamaBahanMakanan string    `json:"nama_bahan_makanan"`
	DeskripsiBahan   string    `json:"deskripsi_bahan"`
}

type BahanMakananPaginationResponse struct {
	Data []BahanMakananResponse `json:"data"`
	PaginationResponse
}

type BahanMakananSimpleResponse struct {
    ID               uuid.UUID `json:"id"`
    NamaBahanMakanan string    `json:"nama_bahan_makanan"`
}