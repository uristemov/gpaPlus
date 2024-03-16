package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/main", h.getMainPage)

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/subject/:id", h.getCoursePageById)

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

	teacher := router.Group("/teach", h.authMiddleware(), h.isTeacher()) // check
	{
		teacher.GET("/courses", h.getAllTeacherCourses)
		teacher.GET("/:id", h.getTeacherById)
		teacher.GET("", h.getAllTeachers)
		teacher.GET(":id/students", h.getAllCourseStudents)
	}

	course := router.Group("/courses", h.authMiddleware())
	{
		course.GET("/:id", h.getCourseById)
		course.GET("", h.getAllCourses)
		//teacher.GET("/search/:name", h.getAuthorByName)
		course.POST("", h.createCourse)
		course.DELETE("/:id", h.deleteCourseById)
		course.PUT("/:id", h.updateCourseById)

		module := router.Group(":course_id/modules")
		{
			module.GET("", h.getAllModuleWithSteps) // done
		}
	}

	module := router.Group("/modules", h.authMiddleware())
	{
		module.GET("/steps/:id", h.getAllModuleSteps) //check
		module.GET("/:id", h.getModuleById)
		module.POST("", h.createModule)
		module.DELETE("/:id", h.deleteModuleById)
		module.PUT("/:id", h.updateModuleById)
	}

	video := router.Group("/videos", h.authMiddleware()) //check
	{
		video.GET("/:id", h.getVideoById)
		video.POST("", h.createVideo)
		video.DELETE("/:id", h.deleteVideoById)
		video.PUT("/:id", h.updateVideoById)
	}

	text := router.Group("/texts", h.authMiddleware()) //check
	{
		text.GET("/:id", h.getTextById)
		text.POST("", h.createText)
		text.DELETE("/:id", h.deleteTextById)
		text.PUT("/:id", h.updateTextById)
	}

	image := router.Group("/images", h.authMiddleware()) //check
	{
		image.GET("/:id", h.getImageById)
		image.POST("", h.createImage)
		image.DELETE("/:id", h.deleteImageById)
		image.PUT("/:id", h.updateImageById)
	}

	requests := router.Group("/requests", h.authMiddleware()) //check
	{
		requests.GET("", h.getAllTeacherRequests)
		requests.POST("", h.createRequest)
		requests.PUT("/:id", h.updateRequestById)
	}

	return router
}
