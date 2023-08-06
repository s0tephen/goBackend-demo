package Indexs

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goBackend-demo/app/request"
	"goBackend-demo/dao/redisServer"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/logUtils"
	"goBackend-demo/utils/services"
	"goBackend-demo/utils/text"
	"goBackend-demo/utils/validateUtils"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Login 用户登录
func Login(ctx *gin.Context) {
	loginReq := request.LoginRequest{}
	loginIp := logUtils.GetRealIP(ctx)
	err := ctx.ShouldBindJSON(&loginReq)
	message, hasErr := validateUtils.ReturnValidateMessage(&loginReq, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}
	userM, tokenM, message, err := logUtils.AuthenticateUser(ctx, &loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New(message, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.New("登陆成功", map[string]interface{}{
		"loginIp":   loginIp,
		"loginTime": time.Now().Format("2006-01-02 15:04:05"),
		"user": gin.H{
			"isAdmin":  userM.IsAdmin,
			"username": userM.Username,
			"token":    tokenM.Token,
		},
	}))
}

// ForgetPwd 找回密码
func ForgetPwd(ctx *gin.Context) {
	forgetPwdReq := request.ForgetPwdRequest{}
	err := ctx.ShouldBindJSON(&forgetPwdReq)
	message, hasErr := validateUtils.ReturnValidateMessage(&forgetPwdReq, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}
	u := dal.User
	userM, err := u.WithContext(ctx).Where(u.Username.Eq(forgetPwdReq.Email)).First()
	if userM == nil {
		ctx.JSON(http.StatusBadRequest, response.New("邮箱不存在", err.Error()))
		return
	}
	//验证邮箱验证码
	code, err := redisServer.Get(forgetPwdReq.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("验证码已过期", err.Error()))
		return
	}
	if forgetPwdReq.EmailCode != code {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码错误", nil))
		return
	}
	//修改密码
	hashPassword, err := services.EncryptPassword(forgetPwdReq.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("修改失败-请联系管理员", err.Error()))
		return
	}
	_, err = u.WithContext(ctx).Where(u.Uemail.Eq(forgetPwdReq.Email)).Update(u.Password, hashPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("修改密码失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.New("修改密码成功", nil))
}

func Login4(ctx *gin.Context) {
	loginReq := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginReq)
	loginIp := logUtils.GetRealIP(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("绑定数据失败", err))
		return
	}
	message, hasErr := validateUtils.ReturnValidateMessage(&loginReq, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}

	u := dal.User

	userM, err := u.WithContext(ctx).Where(u.Username.Eq(loginReq.Username)).First()
	if userM == nil {
		ctx.JSON(http.StatusBadRequest, response.New("用户不存在", err.Error()))
		return
	}

	//获取数据库密码 并对比
	sqlUser, err := u.WithContext(ctx).Where(u.Username.Eq(loginReq.Username)).Select(u.UID, u.Username, u.Password).First()
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(sqlUser.Password), []byte(loginReq.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("密码错误", err.Error()))
		return
	}

	//生成token
	tokenM := model.LoginSession{
		Token:   text.GetUUID(),
		UID:     userM.UID,
		LoginIP: &loginIp,
	}
	loginSession := dal.LoginSession

	//把登陆信息token存入mysql
	loginSession.WithContext(ctx).Create(&tokenM)

	jsonM, _ := json.Marshal(userM)

	//token存入redis
	err = redisServer.Set(fmt.Sprintf("user_token_%s", tokenM.Token), string(jsonM), 1*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("登录失败（token生成失败) 请联系管理员", err.Error()))
		return
	}
	redisServer.PutSet(fmt.Sprintf("user_tokens_%s", userM.Username), []string{tokenM.Token})

	ctx.JSON(http.StatusOK, response.New("登陆成功", map[string]interface{}{
		"loginIp":   loginIp,
		"loginTime": time.Now(),
		"user": gin.H{
			"username": userM.Username,
			"token":    tokenM.Token,
		},
		//"user":      userM,
	}))
}
