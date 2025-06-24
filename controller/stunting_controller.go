package controller

import (
	"net/http"

	"github.com/Revprm/Nutrigrow-Backend/dto"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/Revprm/Nutrigrow-Backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	StuntingController interface {
		Create(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		GetByUserID(ctx *gin.Context)
		GetLatestByUserID(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		Predict(ctx *gin.Context)
		GetForCalendar(ctx *gin.Context)
	}

	stuntingController struct {
		stuntingService service.StuntingService
	}
)

func NewStuntingController(stuntingService service.StuntingService) StuntingController {
	return &stuntingController{
		stuntingService: stuntingService,
	}
}

func (c *stuntingController) Create(ctx *gin.Context) {
	var req dto.StuntingCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userIDStr := ctx.MustGet("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid user ID format", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	req.UserID = userID

	createdStunting, err := c.stuntingService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_STUNTING, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_STUNTING, createdStunting)
	ctx.JSON(http.StatusOK, res)
}

func (c *stuntingController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.stuntingService.GetByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STUNTING, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_STUNTING, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *stuntingController) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	var req dto.PaginationRequest
	// FIX: Changed ShouldBind to ShouldBindQuery for GET requests with query parameters
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.stuntingService.GetByUserID(ctx.Request.Context(), userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STUNTING_BY_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_STUNTING,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *stuntingController) GetLatestByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	result, err := c.stuntingService.GetLatestByUserID(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STUNTING_BY_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_STUNTING, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *stuntingController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.StuntingUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.stuntingService.Update(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_STUNTING, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_STUNTING, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *stuntingController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.stuntingService.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_STUNTING, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_STUNTING, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *stuntingController) Predict(ctx *gin.Context) {
	var req dto.StuntingPredictRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.stuntingService.Predict(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PREDICT_STUNTING, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_PREDICT_STUNTING, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *stuntingController) GetForCalendar(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)

	var req dto.StuntingCalendarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.stuntingService.GetForCalendar(ctx.Request.Context(), userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STUNTING_CALENDAR, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_STUNTING_CALENDAR, result)
	ctx.JSON(http.StatusOK, res)
}
