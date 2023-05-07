package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"index_Demo/docs"
	"index_Demo/gen/response"
	"index_Demo/service/admin_services"
	"index_Demo/service/bing_wallpaper"
	"index_Demo/service/file_services"
	"index_Demo/service/user_services"
	"index_Demo/utils/middleware"
	"index_Demo/utils/middleware/auth"
	"net/http"
)

func Router(g *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.GET("/wallpaper", bing_wallpaper.Wallpaper)

	g.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, response.New("Method not allowed", nil))
	})
	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.New("404 Not found", nil))
	})
	g.Use(middleware.Cors())
	g.Use(middleware.ErrorHandler())
	g.Use(middleware.DeviceType())

	g.GET("/client", middleware.HandleDeviceType)

	g.GET("/post_list", user_services.PostList)
	g.GET("/post", user_services.PostDetail)

	users := g.Group("/user")
	{
		users.POST("/reg_email", user_services.RegEmailCode)
		users.POST("/register", user_services.Register)
		users.POST("/login", user_services.Login)
		users.POST("/update_user_avatar", auth.Middleware(), user_services.UpdateUserAvatar)
		users.POST("/message", auth.Middleware(), user_services.Message)
		users.POST("/feedback", auth.Middleware(), user_services.FeedBack)
		users.POST("/logout", auth.Middleware(), user_services.Logout)

		users.POST("/publish", auth.Middleware(), user_services.PublishPost)
		users.POST("/upload_file", auth.Middleware(), file_services.UploadFile)
	}

	root := g.Group("/root")
	{
		root.GET("/view_user_list", auth.Middleware(), admin_services.ViewUserList)
		root.GET("/review_feedback", auth.Middleware(), admin_services.ReviewFeedback)
	}
}
