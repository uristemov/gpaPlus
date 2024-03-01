package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) getAllTeachers(ctx *gin.Context) {

	teachers, err := h.service.GetAllTeachers(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teachers)
}

func (h *Handler) getTeacherById(ctx *gin.Context) {

	teacherId := ctx.Param("id")
	if teacherId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "teacher id is empty"})
		return
	}

	teacher, err := h.service.GetTeacherById(ctx, teacherId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}
