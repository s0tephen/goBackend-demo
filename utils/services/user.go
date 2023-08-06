package services

import (
	"github.com/gin-gonic/gin"
	"github.com/o1egl/govatar"
	"goBackend-demo/app/request"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/gen/response"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// UpdateUser 用户信息更新
func UpdateUser(ctx *gin.Context, user *model.User, updateRequest request.UpdateRequest) *model.User {
	u := dal.User
	if UserExist(ctx, updateRequest.Username) {
		ctx.JSON(422, response.New("用户名已存在", nil))
		return nil
	}
	hashPassword, err := EncryptPassword(updateRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("注册失败-请联系管理员", err.Error()))
		return nil
	}
	upLoadUser := model.User{
		Username: updateRequest.Username,
		Password: hashPassword,
		CreateAt: time.Now(),
	}
	_, err = u.WithContext(ctx).Where(u.UID.Eq(user.UID)).Updates(upLoadUser)
	if err != nil {
		ctx.JSON(422, response.New("更新失败", nil))
		return nil
	}
	updatedUser, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	return updatedUser
}

// UserExist 判断用户是否存在
func UserExist(ctx *gin.Context, username string) bool {
	u := dal.User
	user, _ := u.WithContext(ctx).Where(u.Username.Eq(username)).First()
	return user != nil
}

// EncryptPassword 加密密码
func EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

// CreateUser 创建用户
func CreateUser(regRequest request.RegRequest, avatar string, hashPassword string) *model.User {
	return &model.User{
		Username: regRequest.Username,
		Avatar:   &avatar,
		Uemail:   &regRequest.Email,
		Password: hashPassword,
		CreateAt: time.Now(),
	}
}

// UserAvatar 生成用户头像
func UserAvatar(uid govatar.Gender, uname string) (string, error) {
	// 创建用户文件夹
	err := os.MkdirAll("./static/images/"+uname+"/avatar", os.ModePerm)
	if err != nil {
		return "", err
	}

	err = govatar.GenerateFileForUsername(uid, uname, "./static/images/"+uname+"/avatar/"+uname+".png")
	if err != nil {
		return "", err
	}

	file, err := os.Open("./static/images/" + uname + "/avatar/" + uname + ".png")
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	// 读取文件内容
	size := fileInfo.Size()
	bytes := make([]byte, size)
	_, err = file.Read(bytes)
	if err != nil {
		return "", err
	}

	return "./static/images/" + uname + "/avatar/" + uname + ".png", nil
}
