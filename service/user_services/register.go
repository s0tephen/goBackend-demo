package user_services

import (
	"github.com/gin-gonic/gin"
	"github.com/o1egl/govatar"
	"golang.org/x/crypto/bcrypt"
	"index_Demo/app/request"
	"index_Demo/dao/redisServer"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"index_Demo/utils/validateUtils"
	"net/http"
	"os"
	"time"
)

// Register 用户注册
func Register(ctx *gin.Context) {

	regRequest := request.RegRequest{}
	err := ctx.BindJSON(&regRequest)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, response.New("系统内部出错 请联系管理员", err.Error()))
	//	return
	//}
	message, hasErr := validateUtils.ReturnValidateMessage(&regRequest, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}

	u := dal.User
	user, _ := u.WithContext(ctx).Where(u.Username.Eq(regRequest.Username)).First()
	if user != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("用户名已存在", nil))
		return
	}

	//从redis中获取验证码
	code, err := redisServer.Get(regRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码已过期", err.Error()))
		return
	}
	if regRequest.EmailCode != code {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码错误", nil))
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(regRequest.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("密码加密失败", err.Error()))
		return
	}
	//avatarBase64, err := generateAvatarBase64(1)
	//if err != nil {
	//	return
	//}
	avatar, err := userAvatar(1, regRequest.Username)
	if err != nil {
		return
	}
	user = &model.User{
		Username: regRequest.Username,
		Avatar:   &avatar,
		Uemail:   &regRequest.Email,
		Password: string(hashPassword),
		CreateAt: time.Now(),
	}

	err = dal.User.Create(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("注册失败", err.Error()))
	}

	//返回给前端的数据
	type UserJson struct {
		Username string    `json:"username"`
		CreatAt  time.Time `json:"creat_at"`
	}
	userJson := UserJson{
		Username: regRequest.Username,
		CreatAt:  time.Now(),
	}
	ctx.JSON(http.StatusOK, response.New("注册成功", userJson))
}

// generateAvatarBase64 生成头像
func userAvatar(uid govatar.Gender, uname string) (string, error) {
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
