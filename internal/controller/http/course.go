package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) getAllCourses(ctx *gin.Context) {

	courses, err := h.service.GetAllCourses(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (h *Handler) getCourseById(ctx *gin.Context) {

	courseId := ctx.Param("id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "course id is empty"})
		return
	}

	course, err := h.service.GetCourseById(ctx, courseId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}
