package controller

import (
	"net/http"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/Revprm/Nutrigrow-Backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	BeritaController interface {
		Create(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		GetByKategori(ctx *gin.Context)
		GetAllWithPagination(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	beritaController struct {
		beritaService service.BeritaService
	}
)

func NewBeritaController(bs service.BeritaService) BeritaController {
	return &beritaController{
		beritaService: bs,
	}
}

func (c *beritaController) Create(ctx *gin.Context) {
	var req dto.BeritaCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.beritaService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_BERITA, result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *beritaController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("ID parameter is required", "missing id parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.beritaService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_BERITA, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *beritaController) GetByKategori(ctx *gin.Context) {
	kategoriID := ctx.Param("kategori_id")
	if kategoriID == "" {
		res := utils.BuildResponseFailed("Kategori ID parameter is required", "missing kategori_id parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.beritaService.GetByKategori(ctx.Request.Context(), kategoriID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LIST_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_LIST_BERITA,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *beritaController) GetAllWithPagination(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Set default values if not provided
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PerPage <= 0 {
		req.PerPage = 10
	}

	result, err := c.beritaService.GetAllWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LIST_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_LIST_BERITA,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *beritaController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("ID parameter is required", "missing id parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.BeritaUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.beritaService.Update(ctx.Request.Context(), req, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_BERITA, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *beritaController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("ID parameter is required", "missing id parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.beritaService.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_BERITA, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_BERITA, nil)
	ctx.JSON(http.StatusOK, res)
}
