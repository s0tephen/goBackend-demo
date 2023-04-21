package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/response"
	"net/http"
)

type LogoutReq struct {
	UID      int32  `json:"uid"`
	Username string `json:"username"`
}

func Logout(ctx *gin.Context) {
	u := dal.User
	logoutReq := new(LogoutReq)
	err := ctx.BindJSON(logoutReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("系统内部出错 请联系管理员", err.Error()))
		return
	}
	result, _ := u.WithContext(ctx).Where(u.UID.Eq(logoutReq.UID), u.Username.Eq(logoutReq.Username)).Delete()
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusServiceUnavailable, response.New("注销失败", nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("注销成功", result))
}
