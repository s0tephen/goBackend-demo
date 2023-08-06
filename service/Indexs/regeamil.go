package Indexs

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goBackend-demo/dao/redisServer"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/email"
	"net/http"
	"time"
)

type RegEmail struct {
	Email string `json:"email" binding:"required"`
}
type ForgetEmail struct {
	Email string `json:"email" binding:"required"`
}

// RegEmailCode 注册邮箱验证码
func RegEmailCode(ctx *gin.Context) {

	regEmail := RegEmail{}
	if err := ctx.BindJSON(&regEmail); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("系统出错联系管理员", err))
		return
	}
	u := dal.User
	if user, _ := u.WithContext(ctx).Where(u.Uemail.Eq(regEmail.Email)).First(); user != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("邮箱已存在", nil))
		return
	}
	if email.VerifyEmailFormat(regEmail.Email) == false {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("邮箱格式错误", nil))
		return
	}
	code, emailErr := email.SendMail(regEmail.Email)
	if emailErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码发送失败 联系系统管理员", emailErr))
		return
	}
	redisServer.Set(regEmail.Email, *code, time.Duration(viper.GetInt("redis.emailTime"))*time.Minute)
	ctx.JSON(http.StatusOK, response.New("验证码发送成功", code))
}

// ForgetCode 忘记密码邮箱验证码
func ForgetCode(ctx *gin.Context) {
	forgetEmail := ForgetEmail{}
	if err := ctx.BindJSON(&forgetEmail); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("系统出错联系管理员", err))
		return
	}
	u := dal.User
	if user, _ := u.WithContext(ctx).Where(u.Uemail.Eq(forgetEmail.Email)).First(); user != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("邮箱已存在", nil))
		return
	}

	code, emailErr := email.SendMail(forgetEmail.Email)
	if emailErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码发送失败 联系系统管理员", emailErr))
		return
	}
	redisServer.Set(forgetEmail.Email, *code, time.Duration(viper.GetInt("redis.emailTime"))*time.Minute)

	ctx.JSON(http.StatusOK, response.New("验证码发送成功", code))
}
