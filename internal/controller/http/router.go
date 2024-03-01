package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", h.createUser)
	router.POST("/refresh", h.refresh)
	router.POST("/login", h.loginUser)

	user := router.Group("/users", h.authMiddleware())
	{
		user.PUT("/", h.updateUser)
		user.DELETE("/", h.deleteUser)
		user.GET("/", h.getUser)
		//user.POST("/confirm", h.confirmUser)
	}

	teacher := router.Group("/teachers")
	{
		teacher.GET("/:id", h.getTeacherById)
		teacher.GET("", h.getAllTeachers)
		//teacher.GET("/search/:name", h.getAuthorByName)
		//teacher.POST("", h.createAuthor)
		//teacher.DELETE("/:id", h.deleteAuthor)
		//teacher.PUT("/:id", h.updateAuthor)
	}

	course := router.Group("/courses")
	{
		course.GET("/:id", h.getCourseById)
		course.GET("", h.getAllCourses)
		//teacher.GET("/search/:name", h.getAuthorByName)
		//teacher.POST("", h.createAuthor)
		//teacher.DELETE("/:id", h.deleteAuthor)
		//teacher.PUT("/:id", h.updateAuthor)
	}

	return router
}
