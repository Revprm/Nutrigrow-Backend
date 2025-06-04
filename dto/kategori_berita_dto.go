package dto

import (
	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_CREATE_KATEGORI_BERITA   = "failed create kategori berita"
	MESSAGE_FAILED_GET_KATEGORI_BERITA      = "failed get kategori berita"
	MESSAGE_FAILED_UPDATE_KATEGORI_BERITA   = "failed update kategori berita"
	MESSAGE_FAILED_DELETE_KATEGORI_BERITA   = "failed delete kategori berita"
	MESSAGE_FAILED_LIST_KATEGORI_BERITA     = "failed get list kategori berita"
	MESSAGE_SUCCESS_CREATE_KATEGORI_BERITA  = "success create kategori berita"
	MESSAGE_SUCCESS_GET_KATEGORI_BERITA     = "success get kategori berita"
	MESSAGE_SUCCESS_UPDATE_KATEGORI_BERITA  = "success update kategori berita"
	MESSAGE_SUCCESS_DELETE_KATEGORI_BERITA  = "success delete kategori berita"
	MESSAGE_SUCCESS_LIST_KATEGORI_BERITA    = "success get list kategori berita"
)

type KategoriBeritaCreateRequest struct {
	NamaKategoriBerita string `json:"nama_kategori_berita" binding:"required"`
}

type KategoriBeritaUpdateRequest struct {
	NamaKategoriBerita string `json:"nama_kategori_berita"`
}

type KategoriBeritaResponse struct {
	ID                 uuid.UUID `json:"id"`
	NamaKategoriBerita string    `json:"nama_kategori_berita"`
}

type KategoriBeritaPaginationResponse struct {
	Data []KategoriBeritaResponse `json:"data"`
	PaginationResponse
}