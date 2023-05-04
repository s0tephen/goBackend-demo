package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/app/request"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"index_Demo/utils/services"
	"index_Demo/utils/validateUtils"
	"net/http"
	"os"
	"time"
)

type UserJson struct {
	Username string    `json:"username"`
	CreatAt  time.Time `json:"creat_at"`
}

// Register 用户注册
func Register(ctx *gin.Context) {
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
	code, err := services.GetCodeFromRedis(regRequest.Email)
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
		ctx.JSON(http.StatusInternalServerError, response.New("密码加密失败", err.Error()))
		return
	}

	avatar, err := services.UserAvatar(1, regRequest.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("生成头像失败", err.Error()))
		return
	}

	user := services.CreateUser(regRequest, avatar, hashPassword)

	err = services.SaveUser(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("注册失败-请联系管理员", err.Error()))
		return
	}

	userJson := UserJson{
		Username: regRequest.Username,
		CreatAt:  time.Now(),
	}
	ctx.JSON(http.StatusOK, response.New("注册成功", userJson))
}

func UpdateUserAvatar(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("无法获取头像文件", err.Error()))
		return
	}
	// 删除旧头像
	oldAvatarPath := "./static/images/" + user.Username + "/avatar/" + user.Username + ".png"
	err = os.Remove(oldAvatarPath)
	if err != nil && !os.IsNotExist(err) {
		ctx.JSON(http.StatusInternalServerError, response.New("系统内部出错", err.Error()))
		return
	}
	// 保存新头像
	avatarPath := "./static/images/" + user.Username + "/avatar/" + user.Username + ".png"
	err = ctx.SaveUploadedFile(file, avatarPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("无法获取头像文件", err.Error()))
		return
	}

	//更新数据库中的头像路径 有点多余
	u := dal.User
	_, err = u.WithContext(ctx).Where(u.Username.Eq(user.Username)).Update(u.Avatar, avatarPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("无法获取头像路径", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.New("头像更新成功", nil))
}
