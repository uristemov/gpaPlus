package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/controller/dto"
	"github.com/uristemov/repeatPro/internal/entity"
	"net/http"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: fmt.Sprintf("invalid input body err: %s", err.Error())})
		return
	}

	userId, err := h.service.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, api.Response{Message: userId})
}

func (h *Handler) loginUser(ctx *gin.Context) {

	var req api.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: fmt.Sprintf("invalid input body err: %s", err.Error())})
		return
	}

	accessToken, refreshToken, err := h.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken, "refreshToken": refreshToken},
	)
}

func (h *Handler) updateUser(ctx *gin.Context) {

	var req api.UpdateUserRequest

	userID, err := getUserId(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	if req == (api.UpdateUserRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{"Update user data not provided"})
		return
	}
	err = h.service.UpdateUser(ctx, userID, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{"User data updated!"})
}

func (h *Handler) getUser(ctx *gin.Context) {

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	user, err := h.service.GetUserById(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, user)
}

func (h *Handler) deleteUser(ctx *gin.Context) {

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	err = h.service.DeleteUser(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{"User deleted"})
}

func (h *Handler) refresh(ctx *gin.Context) {

	var oldRefreshToken dto.RefreshInput

	if err := ctx.BindJSON(&oldRefreshToken); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: fmt.Sprintf("invalid input body err: %s", err.Error())})
		return
	}

	accessToken, refreshToken, err := h.service.Refresh(ctx, oldRefreshToken.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken, "refreshToken": refreshToken},
	)
}
