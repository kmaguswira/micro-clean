package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/controllers"
	"github.com/kmaguswira/micro-clean/app/web/middlewares"
)

func SetupRouter(env string) *gin.Engine {
	var router *gin.Engine

	if env == "local_test" {
		router = gin.Default()
		gin.SetMode(gin.TestMode)
	} else {
		router = gin.New()
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		router.Static("/public", "./public")
	}

	// CORS
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.AuthorizationMiddleware())

	RouterV1(router)

	return router
}

func RouterV1(router *gin.Engine) {
	// setup router
	v1 := router.Group("v1")
	{
		healthGroup := v1.Group("health")
		{
			health := new(controllers.HealthController)

			healthGroup.GET("/check", health.Check)
		}

		// uploadGroup := v1.Group("upload")
		// {
		// 	upload := new(controllers.UploadController)

		// 	uploadGroup.POST("/picture", upload.Picture)
		// }

		authGroup := v1.Group("auth")
		{
			auth := new(controllers.AuthController)
			// user := new(controllers.UserController)

			authGroup.POST("/sign-up", auth.SignUp)
			authGroup.POST("/sign-in", auth.SignIn)
			// authGroup.GET("/self", auth.Self)
			// authGroup.POST("/sign-in/admin", auth.PostSignInAdmin)
			// authGroup.GET("/all", user.GetUsers)

		}
	}

}
