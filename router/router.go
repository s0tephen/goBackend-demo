package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"index_Demo/docs"
	"index_Demo/service/bing_wallpaper"
	"index_Demo/service/user_services"
	"index_Demo/utils/middleware"
	"index_Demo/utils/middleware/auth"
)

func Router(g *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.GET("/wallpaper", bing_wallpaper.Wallpaper)

	g.Use(middleware.Cors())
	g.Use(middleware.ErrorHandler())

	users := g.Group("/user")

	{
		users.POST("/register", user_services.Register)
		users.POST("/login", user_services.Login)
		users.POST("/logout", auth.Middleware(), user_services.Logout)
		users.POST("/message", auth.Middleware(), user_services.Message)
	}
}
