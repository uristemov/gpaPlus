package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) createCourse(ctx *gin.Context) {

	var req api.CreateCourseRequest

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

	courseId, err := h.service.CreateCourse(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: courseId})
}

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

func (h *Handler) updateCourseById(ctx *gin.Context) {

	var req api.UpdateCourseRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	courseId := ctx.Param("id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Module id is empty"})
		return
	}

	if req == (api.UpdateCourseRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
		return
	}

	err := h.service.UpdateCourseById(ctx, &req, courseId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteCourseById(ctx *gin.Context) {

	courseId := ctx.Param("id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Course id is empty"})
		return
	}

	err := h.service.DeleteCourseById(ctx, courseId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "Course is deleted"})
}

func (h *Handler) getAllTeacherCourses(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	courses, err := h.service.GetAllTeacherCourses(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (h *Handler) getAllCourseStudents(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	courseId := ctx.Param("id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Course id is empty"})
		return
	}

	students, err := h.service.GetAllCourseStudents(ctx, userId, courseId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}
