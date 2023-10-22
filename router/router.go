package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goBackend-demo/docs"
	"goBackend-demo/gen/response"
	"goBackend-demo/service/Indexs"
	"goBackend-demo/service/admin"
	"goBackend-demo/service/bing_wallpaper"
	"goBackend-demo/service/file"
	"goBackend-demo/service/user"
	"goBackend-demo/service/user/updates"
	"goBackend-demo/utils/middleware"
	"goBackend-demo/utils/middleware/auth"
	"net/http"
)

func Router(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/wallpaper", bing_wallpaper.Wallpaper)
	router.Static("/static", "./static")

	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, response.New("Method not allowed", nil))
	})
	//router.NoRoute(func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusNotFound, response.New("404 Not found", nil))
	//})
	router.Use(middleware.Cors())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.DeviceType())

	router.GET("/client", middleware.HandleDeviceType)
	router.GET("/post_list", user.PostList)
	router.GET("/post", user.PostDetail)

	apiRouter := router.Group("/api")
	//limitReq := middleware.NewRateLimiter()
	//apiRouter.Use(limitReq.LimitRequest(3, 10*time.Minute))
	// index
	Index := apiRouter.Group("/")
	{
		Index.POST("/reg_email", Indexs.RegEmailCode)
		Index.POST("/register", Indexs.Register)
		Index.POST("/login", Indexs.Login)

		//Forget
		Index.POST("/forget_email", Indexs.ForgetCode)
		Index.POST("/forget", Indexs.ForgetPwd)
	}

	// auth
	auths := apiRouter.Group("/auth")
	auths.Use(auth.Middleware())
	{
		auths.POST("/update", updates.UpdateUserInfo)
		auths.POST("/update_avatar", updates.UpdateUserAvatar)
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
		root.GET("/userlist", admin.ViewUserList)
		root.GET("/delete_user", admin.DeleteUser)

		root.GET("/review_feedback", admin.ReviewFeedback)
	}
}
