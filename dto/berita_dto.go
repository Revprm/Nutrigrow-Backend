package dto

import (
	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_CREATE_BERITA   = "failed create berita"
	MESSAGE_FAILED_GET_BERITA      = "failed get berita"
	MESSAGE_FAILED_UPDATE_BERITA   = "failed update berita"
	MESSAGE_FAILED_DELETE_BERITA   = "failed delete berita"
	MESSAGE_FAILED_LIST_BERITA     = "failed get list berita"
	MESSAGE_SUCCESS_CREATE_BERITA  = "success create berita"
	MESSAGE_SUCCESS_GET_BERITA     = "success get berita"
	MESSAGE_SUCCESS_UPDATE_BERITA  = "success update berita"
	MESSAGE_SUCCESS_DELETE_BERITA  = "success delete berita"
	MESSAGE_SUCCESS_LIST_BERITA    = "success get list berita"
)

type BeritaCreateRequest struct {
	JudulBerita      string    `json:"judul_berita" binding:"required"`
	SubjudulBerita   string    `json:"subjudul_berita"`
	DeskripsiBerita  string    `json:"deskripsi_berita" binding:"required"`
	KategoriBeritaID uuid.UUID `json:"kategori_berita_id" binding:"required"`
}

type BeritaUpdateRequest struct {
	JudulBerita      string    `json:"judul_berita"`
	SubjudulBerita   string    `json:"subjudul_berita"`
	DeskripsiBerita  string    `json:"deskripsi_berita"`
	KategoriBeritaID uuid.UUID `json:"kategori_berita_id"`
}

type BeritaResponse struct {
	ID               uuid.UUID             `json:"id"`
	JudulBerita      string                `json:"judul_berita"`
	SubjudulBerita   string                `json:"subjudul_berita"`
	DeskripsiBerita  string                `json:"deskripsi_berita"`
	KategoriBeritaID uuid.UUID             `json:"kategori_berita_id"`
	KategoriBerita   KategoriBeritaResponse `json:"kategori_berita,omitempty"`
}

type BeritaPaginationResponse struct {
	Data []BeritaResponse `json:"data"`
	PaginationResponse
}