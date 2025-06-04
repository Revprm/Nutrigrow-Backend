package controller

import (
	"net/http"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/Revprm/Nutrigrow-Backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	MakananController interface {
		Create(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		GetByNama(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetByBahanMakanan(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	makananController struct {
		makananService service.MakananService
	}
)

func NewMakananController(ms service.MakananService) MakananController {
	return &makananController{
		makananService: ms,
	}
}

func (c *makananController) Create(ctx *gin.Context) {
	var req dto.MakananCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_MAKANAN_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.makananService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *makananController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.makananService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *makananController) GetByNama(ctx *gin.Context) {
	nama := ctx.Param("nama")

	result, err := c.makananService.GetByNama(ctx.Request.Context(), nama)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *makananController) GetAll(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_MAKANAN_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.makananService.GetAllWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LIST_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LIST_MAKANAN, result.Data)
	res.Meta = result.PaginationResponse
	ctx.JSON(http.StatusOK, res)
}

func (c *makananController) GetByBahanMakanan(ctx *gin.Context) {
	bahanID := ctx.Param("bahanId")

	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_MAKANAN_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.makananService.GetByBahanMakanan(ctx.Request.Context(), bahanID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LIST_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LIST_MAKANAN, result.Data)
	res.Meta = result.PaginationResponse
	ctx.JSON(http.StatusOK, res)
}

func (c *makananController) Update(ctx *gin.Context) {
	var req dto.MakananUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_MAKANAN_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id := ctx.Param("id")

	result, err := c.makananService.Update(ctx.Request.Context(), req, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_MAKANAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_MAKANAN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *makananController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.makananService.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_MAKANAN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_MAKANAN, nil)
	ctx.JSON(http.StatusOK, res)
}