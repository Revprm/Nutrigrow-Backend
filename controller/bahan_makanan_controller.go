package controller

import (
	"net/http"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/Revprm/Nutrigrow-Backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	BahanMakananController interface {
		Create(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		GetByNama(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	bahanMakananController struct {
		bahanMakananService service.BahanMakananService
	}
)

func NewBahanMakananController(bs service.BahanMakananService) BahanMakananController {
	return &bahanMakananController{
		bahanMakananService: bs,
	}
}

func (c *bahanMakananController) Create(ctx *gin.Context) {
	var req dto.BahanMakananCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.bahanMakananService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BAHAN_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_BAHAN_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bahanMakananController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.bahanMakananService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BAHAN_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_BAHAN_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bahanMakananController) GetByNama(ctx *gin.Context) {
	nama := ctx.Param("nama")

	result, err := c.bahanMakananService.GetByNama(ctx.Request.Context(), nama)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BAHAN_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_BAHAN_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bahanMakananController) GetAll(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.bahanMakananService.GetAllWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LIST_BAHAN_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LIST_BAHAN_MAKANAN, result.Data)
	res.Meta = result.PaginationResponse
	ctx.JSON(http.StatusOK, res)
}

func (c *bahanMakananController) Update(ctx *gin.Context) {
	var req dto.BahanMakananUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id := ctx.Param("id")

	result, err := c.bahanMakananService.Update(ctx.Request.Context(), req, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_BAHAN_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_BAHAN_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bahanMakananController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.bahanMakananService.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_BAHAN_MAKANAN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_BAHAN_MAKANAN, nil)
	ctx.JSON(http.StatusOK, res)
}