package admin_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/utils/middleware/auth"
)

func ReviewFeedback(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	user.UID = 1
}
