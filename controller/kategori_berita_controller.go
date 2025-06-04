package controller

import (
	"net/http"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/Revprm/Nutrigrow-Backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	KategoriBeritaController interface {
		Create(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		GetByNama(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	kategoriBeritaController struct {
		kategoriBeritaService service.KategoriBeritaService
	}
)

func NewKategoriBeritaController(kbs service.KategoriBeritaService) KategoriBeritaController {
	return &kategoriBeritaController{
		kategoriBeritaService: kbs,
	}
}

func (c *kategoriBeritaController) Create(ctx *gin.Context) {
	var req dto.KategoriBeritaCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.kategoriBeritaService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_KATEGORI_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_KATEGORI_BERITA, result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *kategoriBeritaController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("ID parameter is required", "missing id parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.kategoriBeritaService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_KATEGORI_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_KATEGORI_BERITA, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *kategoriBeritaController) GetByNama(ctx *gin.Context) {
	nama := ctx.Query("nama")
	if nama == "" {
		res := utils.BuildResponseFailed("Nama parameter is required", "missing nama parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.kategoriBeritaService.GetByNama(ctx.Request.Context(), nama)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_KATEGORI_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_KATEGORI_BERITA, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *kategoriBeritaController) GetAll(ctx *gin.Context) {
	result, err := c.kategoriBeritaService.GetAll(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LIST_KATEGORI_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LIST_KATEGORI_BERITA, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *kategoriBeritaController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("ID parameter is required", "missing id parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.KategoriBeritaUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.kategoriBeritaService.Update(ctx.Request.Context(), req, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_KATEGORI_BERITA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_KATEGORI_BERITA, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *kategoriBeritaController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("ID parameter is required", "missing id parameter", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.kategoriBeritaService.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_KATEGORI_BERITA, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_KATEGORI_BERITA, nil)
	ctx.JSON(http.StatusOK, res)
}
