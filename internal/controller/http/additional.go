package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) getMainPage(ctx *gin.Context) {

	courses, teachers, err := h.service.GetCoursesAndTeachers(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"courses":  courses,
		"teachers": teachers,
	})
}

func (h *Handler) getCourseDataPage(ctx *gin.Context) {

	courseId := ctx.Param("id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Course id is empty"})
		return
	}

	modules, err := h.service.GetAllCourseModules(ctx, courseId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "course.html", gin.H{
		"modules": modules,
	})
}

func (h *Handler) getCoursePageById(ctx *gin.Context) {

	courseId := ctx.Param("id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Course id is empty"})
		return
	}

	course, err := h.service.GetCourseById(ctx, courseId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	fmt.Println(course)

	ctx.HTML(http.StatusOK, "course.html", gin.H{
		"course": course,
	})
}
