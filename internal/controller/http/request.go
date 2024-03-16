package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) createRequest(ctx *gin.Context) {

	var req api.CreateRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	userId, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	req.UserId = userId
	req.IsActive = true

	requestId, err := h.service.CreateRequest(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: requestId})
}

func (h *Handler) getAllTeacherRequests(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	requests, err := h.service.GetAllTeacherRequests(ctx, userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, requests)
}

func (h *Handler) updateRequestById(ctx *gin.Context) {

	var req api.UpdateRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	requestId := ctx.Param("id")
	if requestId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Request id is empty"})
		return
	}

	if req == (api.UpdateRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
		return
	}

	err := h.service.UpdateRequestById(ctx, &req, requestId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}
