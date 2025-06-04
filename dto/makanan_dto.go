package dto

import (
	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_CREATE_MAKANAN   = "failed create makanan"
	MESSAGE_FAILED_GET_MAKANAN      = "failed get makanan"
	MESSAGE_FAILED_GET_MAKANAN_FROM_BODY = "failed get makanan from body"
	MESSAGE_FAILED_UPDATE_MAKANAN   = "failed update makanan"
	MESSAGE_FAILED_DELETE_MAKANAN   = "failed delete makanan"
	MESSAGE_FAILED_LIST_MAKANAN     = "failed get list makanan"
	MESSAGE_SUCCESS_CREATE_MAKANAN  = "success create makanan"
	MESSAGE_SUCCESS_GET_MAKANAN     = "success get makanan"
	MESSAGE_SUCCESS_UPDATE_MAKANAN  = "success update makanan"
	MESSAGE_SUCCESS_DELETE_MAKANAN  = "success delete makanan"
	MESSAGE_SUCCESS_LIST_MAKANAN    = "success get list makanan"
)

type MakananCreateRequest struct {
	NamaMakanan          string      `json:"nama_makanan" binding:"required"`
	DeskripsiMakanan     string      `json:"deskripsi_makanan"`
	VideoTutorialMakanan string      `json:"video_tutorial_makanan"`
	BahanMakananIDs      []uuid.UUID `json:"bahan_makanan_ids"`
}

type MakananUpdateRequest struct {
	NamaMakanan          string      `json:"nama_makanan"`
	DeskripsiMakanan     string      `json:"deskripsi_makanan"`
	VideoTutorialMakanan string      `json:"video_tutorial_makanan"`
	BahanMakananIDs      []uuid.UUID `json:"bahan_makanan_ids"`
}

type MakananResponse struct {
    ID                   uuid.UUID                   `json:"id"`
    NamaMakanan          string                     `json:"nama_makanan"`
    DeskripsiMakanan     string                     `json:"deskripsi_makanan"`
    VideoTutorialMakanan string                     `json:"video_tutorial_makanan"`
    BahanMakanans        []BahanMakananSimpleResponse `json:"bahan_makanans,omitempty"`
}

type MakananPaginationResponse struct {
	Data []MakananResponse `json:"data"`
	PaginationResponse
}