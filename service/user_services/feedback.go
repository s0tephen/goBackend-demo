package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/utils/middleware/auth"
)

func FeedBack(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	if user == nil {

	}
}
