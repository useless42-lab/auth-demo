package routes

import (
	"UserCenter/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func InitApiRoute() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
			auth.POST("/register", controller.Register)
			password := auth.Group("/password")
			{
				password.POST("/change", controller.ChangePassword)
				password.POST("/reset", controller.ResetPassword)
			}
		}
	}
	router.Run(":" + os.Getenv("ROUTE_PORT"))
}
