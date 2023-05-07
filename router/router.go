package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"index_Demo/docs"
	"index_Demo/gen/response"
	"index_Demo/service/admin"
	"index_Demo/service/bing_wallpaper"
	"index_Demo/service/file_services"
	"index_Demo/service/user"
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

	g.GET("/post_list", user.PostList)
	g.GET("/post", user.PostDetail)

	users := g.Group("/user")
	{
		users.POST("/reg_email", user.RegEmailCode)
		users.POST("/register", user.Register)
		users.POST("/login", user.Login)
		users.POST("/update_user_avatar", auth.Middleware(), user.UpdateUserAvatar)
		users.POST("/message", auth.Middleware(), user.Message)
		users.POST("/feedback", auth.Middleware(), user.FeedBack)
		users.POST("/logout", auth.Middleware(), user.Logout)

		users.POST("/publish", auth.Middleware(), user.PublishPost)
		users.POST("/upload_file", auth.Middleware(), file_services.UploadFile)
	}

	root := g.Group("/root")

	{
		root.GET("/view_user_list", auth.Middleware(), admin.ViewUserList)
		root.GET("/review_feedback", auth.Middleware(), admin.ReviewFeedback)
	}
}
