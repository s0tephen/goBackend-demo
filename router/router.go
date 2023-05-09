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

func Router(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/wallpaper", bing_wallpaper.Wallpaper)

	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, response.New("Method not allowed", nil))
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.New("404 Not found", nil))
	})
	router.Use(middleware.Cors())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.DeviceType())

	router.GET("/client", middleware.HandleDeviceType)
	router.GET("/post_list", user.PostList)
	router.GET("/post", user.PostDetail)

	apiRouter := router.Group("/api")

	Index := apiRouter.Group("/")
	{
		Index.POST("/reg_email", user.RegEmailCode)
		Index.POST("/register", user.Register)
	}
	{
		auths := apiRouter.Group("/auth")
		{
			auths.POST("/login", user.Login)
			auths.POST("/update_user_avatar", auth.Middleware(), user.UpdateUserAvatar)
			auths.POST("/message", auth.Middleware(), user.Message)
			auths.POST("/feedback", auth.Middleware(), user.FeedBack)
			auths.POST("/logout", auth.Middleware(), user.Logout)

			auths.POST("/publish", auth.Middleware(), user.PublishPost)
			auths.POST("/upload_file", auth.Middleware(), file_services.UploadFile)
		}

	}

	root := router.Group("/admin")

	{
		root.GET("/view_user_list", auth.Middleware(), admin.ViewUserList)
		root.GET("/review_feedback", auth.Middleware(), admin.ReviewFeedback)
	}
}
