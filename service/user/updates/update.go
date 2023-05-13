package updates

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
)

// UpdateUserInfo 更新用户账号和密码
func UpdateUserInfo(ctx *gin.Context) {
	currentUser := auth.CurrentUser(ctx)
	updateRequest := request.UpdateRequest{}
	err := ctx.BindJSON(&updateRequest)
	message, hasErr := validateUtils.ReturnValidateMessage(&updateRequest, err)
	if hasErr && updateRequest.Username != currentUser.Username {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}
	updatedUser := services.UpdateUser(ctx, currentUser, updateRequest)

	userJson := request.Json{
		Username: updatedUser.Username,
		UploadAt: updatedUser.CreateAt,
	}
	ctx.JSON(http.StatusOK, response.New("更新成功", userJson))
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
	if err = os.Remove(oldAvatarPath); err != nil && !os.IsNotExist(err) {
		ctx.JSON(http.StatusInternalServerError, response.New("系统内部出错", err.Error()))
		return
	}
	// 保存新头像
	avatarPath := "./static/images/" + user.Username + "/avatar/" + user.Username + `.png`
	if err = ctx.SaveUploadedFile(file, avatarPath); err != nil {
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
