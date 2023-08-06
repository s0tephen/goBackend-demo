package user

import (
	"github.com/gin-gonic/gin"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/middleware/auth"
	"net/http"
)

// Logout 注销
func Logout(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	u := dal.User
	result, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID), u.Username.Eq(user.Username)).Delete()
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusServiceUnavailable, response.New("注销失败 请联系管理员", nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("注销成功", user.Username))
}
