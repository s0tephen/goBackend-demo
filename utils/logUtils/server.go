package logUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goBackend-demo/app/request"
	"goBackend-demo/dao/redisServer"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/utils/text"
	"goBackend-demo/utils/validateUtils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetRealIP(ctx *gin.Context) string {
	return ctx.ClientIP()
	//ctx.RemoteIP()
}

// AuthenticateUser 验证Log
func AuthenticateUser(ctx *gin.Context, loginReq *request.LoginRequest) (*model.User, *model.LoginSession, string, error) {
	u := dal.User
	message, hasErr := validateUtils.ReturnValidateMessage(loginReq, nil)
	if hasErr {
		return nil, nil, message, errors.New("无效的登陆请求")
	}
	userM, err := u.WithContext(ctx).Where(u.Username.Eq(loginReq.Username)).Or(u.Uemail.Eq(loginReq.Username)).First()
	if userM == nil {
		return nil, nil, "用户不存在", errors.New("用户不存在")
	}

	// 对比数据库中的密码和请求提供的密码
	sqlUser, err := u.WithContext(ctx).Where(u.Username.Eq(loginReq.Username)).Or(u.Uemail.Eq(loginReq.Username)).Select(u.UID, u.Username, u.Password).First()
	if err != nil {
		return nil, nil, "查询用户失败", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(sqlUser.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, nil, "密码错误", errors.New("无效的登陆凭据")
	}

	loginIp := GetRealIP(ctx)
	Time := time.Now()
	tokenM := &model.LoginSession{
		Token:     text.GetUUID(),
		LoginTime: &Time,
		UID:       userM.UID,
		LoginIP:   &loginIp,
	}

	loginSession := dal.LoginSession

	// 把令牌存入mysql数据库
	loginSession.WithContext(ctx).Create(tokenM)

	// 将用户的令牌存入Redis缓存
	jsonM, _ := json.Marshal(userM)
	err = redisServer.Set(fmt.Sprintf("user_token_%s", tokenM.Token), string(jsonM), 1*time.Hour)
	if err != nil {
		return nil, nil, "登录失败（token生成失败) 请联系管理员", err
	}
	redisServer.PutSet(fmt.Sprintf("user_tokens_%s", userM.Username), []string{tokenM.Token})

	return userM, tokenM, "", nil
}
