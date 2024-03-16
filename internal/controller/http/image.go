package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) createImage(ctx *gin.Context) {

	var req api.CreateImageRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	imageId, err := h.service.CreateImage(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: imageId})
}

func (h *Handler) getImageById(ctx *gin.Context) {

	imageId := ctx.Param("id")
	if imageId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Image id is empty"})
		return
	}

	image, err := h.service.GetImageById(ctx, imageId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, image)
}

func (h *Handler) updateImageById(ctx *gin.Context) {

	var req api.UpdateImageRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	imageId := ctx.Param("id")
	if imageId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Image id is empty"})
		return
	}

	if req == (api.UpdateImageRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
		return
	}

	err := h.service.UpdateImageById(ctx, &req, imageId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteImageById(ctx *gin.Context) {

	imageId := ctx.Param("id")
	if imageId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Image id is empty"})
		return
	}

	err := h.service.DeleteImageById(ctx, imageId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "Image is deleted"})
}
