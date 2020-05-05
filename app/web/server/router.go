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

			authGroup.POST("/sign-up", auth.SignUp)
			authGroup.POST("/sign-in", auth.SignIn)
			authGroup.GET("/activate-user/:token", auth.ActivateUser)
			authGroup.POST("/forgot-password", auth.ForgotPassword)
			authGroup.POST("/reset-password/:token", auth.ResetPassword)
			authGroup.POST("/change-password", auth.ChangePassword)
			authGroup.GET("/self", auth.Self)
		}

		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)

			userGroup.POST("/create", user.Create)
			userGroup.GET("/:id", user.FindById)
			userGroup.DELETE("/:id", user.Delete)
			userGroup.PUT("/:id", user.Update)
			userGroup.GET("/", user.FindAll)
		}

		roleGroup := v1.Group("role")
		{
			role := new(controllers.RoleController)

			roleGroup.POST("/create", role.Create)
			roleGroup.GET("/:id", role.FindById)
			roleGroup.DELETE("/:id", role.Delete)
			roleGroup.PUT("/:id", role.Update)
			roleGroup.GET("/", role.FindAll)
		}

		aclGroup := v1.Group("acl")
		{
			acl := new(controllers.ACLController)

			aclGroup.POST("/create", acl.Create)
			aclGroup.GET("/:id", acl.FindById)
			aclGroup.DELETE("/:id", acl.Delete)
			aclGroup.PUT("/:id", acl.Update)
			aclGroup.GET("/", acl.FindAll)
		}

		emailTemplateGroup := v1.Group("email-template")
		{
			emailTemplate := new(controllers.EmailTemplateController)

			emailTemplateGroup.POST("/create", emailTemplate.Create)
			emailTemplateGroup.GET("/:id", emailTemplate.FindById)
			emailTemplateGroup.DELETE("/:id", emailTemplate.Delete)
			emailTemplateGroup.PUT("/:id", emailTemplate.Update)
			emailTemplateGroup.GET("/", emailTemplate.FindAll)
			emailTemplateGroup.POST("/send-email", emailTemplate.SendEmail)
		}
	}

}
