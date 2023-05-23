package Indexs

import (
	"github.com/gin-gonic/gin"
	"index_Demo/app/request"
	"index_Demo/dao/redisServer"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/response"
	"index_Demo/utils/services"
	"index_Demo/utils/validateUtils"
	"net/http"
	"time"
)

func Register(ctx *gin.Context) {
	CreateUser(ctx)
}

// CreateUser 用户注册
func CreateUser(ctx *gin.Context) {
	regRequest := request.RegRequest{}
	err := ctx.BindJSON(&regRequest)
	message, hasErr := validateUtils.ReturnValidateMessage(&regRequest, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}
	if services.UserExist(ctx, regRequest.Username) {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("用户名已存在", nil))
		return
	}
	code, err := redisServer.Get(regRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码已过期", err.Error()))
		return
	}
	if regRequest.EmailCode != code {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码错误", nil))
		return
	}
	hashPassword, err := services.EncryptPassword(regRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("注册失败-请联系管理员", err.Error()))
		return
	}
	avatar, err := services.UserAvatar(1, regRequest.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("注册失败-请联系管理员", err.Error()))
		return
	}
	user := services.CreateUser(regRequest, avatar, hashPassword)
	if err = dal.User.Create(user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("注册失败-请联系管理员", err.Error()))
		return
	}
	userJson := request.Json{
		Username: regRequest.Username,
		CreatAt:  time.Now(),
	}
	ctx.JSON(http.StatusOK, response.New("注册成功", userJson))
}
