package admin_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"net/http"
)

func User(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	u := dal.User

	//查询数据库该用户对应的信息
	userInfo, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	if userInfo.UID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("用户不存在", nil))
		return
	}
}
