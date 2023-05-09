package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"index_Demo/docs"
	"index_Demo/gen/response"
	"index_Demo/service/admin"
	"index_Demo/service/bing_wallpaper"
	"index_Demo/service/file"
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

	// index
	Index := apiRouter.Group("/")
	{
		Index.POST("/reg_email", user.RegEmailCode)
		Index.POST("/register", user.Register)
		Index.POST("/login", user.Login)
	}

	// auth
	auths := apiRouter.Group("/auth")
	auths.Use(auth.Middleware())
	{
		auths.POST("/update_user_avatar", user.UpdateUserAvatar)
		auths.POST("/message", user.Message)
		auths.POST("/feedback", user.FeedBack)
		auths.POST("/logout", user.Logout)

		auths.POST("/publish", user.PublishPost)
		auths.POST("/upload_file", file.UploadFile)
	}

	// admin
	root := router.Group("/admin")
	root.Use(auth.Middleware())
	{
		root.GET("/view_user_list", admin.ViewUserList)
		root.GET("/review_feedback", admin.ReviewFeedback)
	}
}
