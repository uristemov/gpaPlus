package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) createVideo(ctx *gin.Context) {

	var req api.CreateVideoRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	videoId, err := h.service.CreateVideo(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: videoId})
}

func (h *Handler) getVideoById(ctx *gin.Context) {

	videoId := ctx.Param("id")
	if videoId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "video id is empty"})
		return
	}

	video, err := h.service.GetVideoById(ctx, videoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, video)
}

func (h *Handler) updateVideoById(ctx *gin.Context) {

	var req api.UpdateVideoRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	videoId := ctx.Param("id")
	if videoId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Video id is empty"})
		return
	}

	if req == (api.UpdateVideoRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
		return
	}

	err := h.service.UpdateVideoById(ctx, &req, videoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteVideoById(ctx *gin.Context) {

	videoId := ctx.Param("id")
	if videoId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Video id is empty"})
		return
	}

	err := h.service.DeleteVideoById(ctx, videoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "Video is deleted"})
}
