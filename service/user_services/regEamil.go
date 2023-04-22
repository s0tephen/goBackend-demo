package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/dao/redisServer"
	"index_Demo/gen/response"
	"index_Demo/utils/email"
	"net/http"
	"time"
)

type RegEmail struct {
	Email string `json:"email" binding:"required"`
}

func RegEmailCode(ctx *gin.Context) {
	regEmail := RegEmail{}
	err := ctx.BindJSON(&regEmail)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("系统出错联系管理员", err))
		return
	}
	code, emailErr := email.SendMail(regEmail.Email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码发送失败 联系系统管理员", emailErr))
		return
	}
	redisServer.Set(regEmail.Email, *code, 5*time.Minute)

	ctx.JSON(http.StatusOK, response.New("验证码发送成功", code))
}
