package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"net/http"
)

func (h *Handler) createModule(ctx *gin.Context) {

	var req api.CreateModuleRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}
	fmt.Println("Request ", req)
	moduleId, err := h.service.CreateModule(ctx, &req)
	if err != nil {
		fmt.Println("Request ", req)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: moduleId})
}

func (h *Handler) getAllCourseModules(ctx *gin.Context) {

	courseId := ctx.Param("course_id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "course id is empty"})
		return
	}

	modules, err := h.service.GetAllCourseModules(ctx, courseId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, modules)
}

func (h *Handler) getModuleById(ctx *gin.Context) {

	moduleId := ctx.Param("id")
	if moduleId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "module id is empty"})
		return
	}

	course, err := h.service.GetModuleById(ctx, moduleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func (h *Handler) updateModuleById(ctx *gin.Context) {

	var req api.UpdateModuleRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	moduleId := ctx.Param("id")
	if moduleId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Module id is empty"})
		return
	}

	if req == (api.UpdateModuleRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
		return
	}

	err := h.service.UpdateModuleById(ctx, &req, moduleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteModuleById(ctx *gin.Context) {

	moduleId := ctx.Param("id")
	if moduleId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "module id is empty"})
		return
	}

	err := h.service.DeleteModuleById(ctx, moduleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "Module deleted"})
}

func (h *Handler) getAllModuleSteps(ctx *gin.Context) {

	moduleId := ctx.Param("id")
	if moduleId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "module id is empty"})
		return
	}

	steps, err := h.service.GetAllModuleSteps(ctx, moduleId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, steps)
}

func (h *Handler) getAllModuleWithSteps(ctx *gin.Context) {

	courseId := ctx.Param("course_id")
	if courseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "course id is empty"})
		return
	}

	steps, err := h.service.GetAllModuleWithSteps(ctx, courseId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, steps)
}
