package user

import (
	"github.com/gin-gonic/gin"
	"goBackend-demo/app/request"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/middleware/auth"
	"goBackend-demo/utils/validateUtils"
	"net/http"
	"time"
)

// Message 用户留言
func Message(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	msgRes := request.MessageRes{}
	err := ctx.BindJSON(&msgRes)
	message, hasErr := validateUtils.ReturnValidateMessage(&msgRes, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}
	contentReq := model.Message{
		Uname:    &user.Username,
		CreateAt: time.Now(),
		Content:  &msgRes.Content,
	}
	if dal.Message.Create(&contentReq) != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("留言失败", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.New("留言成功", msgRes.Content))
}
