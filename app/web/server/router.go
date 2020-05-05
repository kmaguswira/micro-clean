package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/micro-clean/app/web/controllers"
	"github.com/kmaguswira/micro-clean/app/web/middlewares"
	"github.com/micro/go-micro/client"
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

	corsMiddleware := middlewares.NewCORSMiddleware()
	router.Use(corsMiddleware.Handler())
	authorizationMiddleware := middlewares.NewAuthorizationMiddleware(client.DefaultClient)
	router.Use(authorizationMiddleware.Handler())

	RouterV1(router)

	return router
}

func RouterV1(router *gin.Engine) {
	// setup router
	v1 := router.Group("v1")
	{
		healthGroup := v1.Group("health")
		{
			health := controllers.NewHealthController()

			healthGroup.GET("/check", health.Check)
		}

		uploadGroup := v1.Group("upload")
		{
			upload := controllers.NewUploadController(client.DefaultClient)

			uploadGroup.POST("/images", upload.Images)
		}

		authGroup := v1.Group("auth")
		{
			auth := controllers.NewAuthController(client.DefaultClient)

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
			user := controllers.NewUserController(client.DefaultClient)

			userGroup.POST("/create", user.Create)
			userGroup.GET("/:id", user.FindById)
			userGroup.DELETE("/:id", user.Delete)
			userGroup.PUT("/:id", user.Update)
			userGroup.GET("/", user.FindAll)
		}

		roleGroup := v1.Group("role")
		{
			role := controllers.NewRoleController(client.DefaultClient)

			roleGroup.POST("/create", role.Create)
			roleGroup.GET("/:id", role.FindById)
			roleGroup.DELETE("/:id", role.Delete)
			roleGroup.PUT("/:id", role.Update)
			roleGroup.GET("/", role.FindAll)
		}

		aclGroup := v1.Group("acl")
		{
			acl := controllers.NewACLController(client.DefaultClient)

			aclGroup.POST("/create", acl.Create)
			aclGroup.GET("/:id", acl.FindById)
			aclGroup.DELETE("/:id", acl.Delete)
			aclGroup.PUT("/:id", acl.Update)
			aclGroup.GET("/", acl.FindAll)
		}

		emailTemplateGroup := v1.Group("email-template")
		{
			emailTemplate := controllers.NewEmailTemplateController(client.DefaultClient)

			emailTemplateGroup.POST("/create", emailTemplate.Create)
			emailTemplateGroup.GET("/:id", emailTemplate.FindById)
			emailTemplateGroup.DELETE("/:id", emailTemplate.Delete)
			emailTemplateGroup.PUT("/:id", emailTemplate.Update)
			emailTemplateGroup.GET("/", emailTemplate.FindAll)
			emailTemplateGroup.POST("/send-email", emailTemplate.SendEmail)
		}
	}

}
