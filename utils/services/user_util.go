package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"index_Demo/app/request"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"net/http"
	"time"
)

func IsLogin(ctx *gin.Context) bool {
	token := ctx.GetHeader("Authorization")
	_, err := auth.GetUserByToken(token)
	if token == "" && err != nil {
		return false
	}
	return true
}

// IsAdmin 管理员权限验证
func IsAdmin(ctx *gin.Context) bool {
	user := auth.CurrentUser(ctx)
	u := dal.User
	userInfo, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	return userInfo.IsAdmin
}

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
		UploadAt: time.Now(),
	}
	_, err = u.WithContext(ctx).Where(u.UID.Eq(user.UID)).Updates(upLoadUser)
	if err != nil {
		ctx.JSON(422, response.New("更新失败", nil))
		return nil
	}
	updatedUser, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	return updatedUser
}

// Query 分页查询
func Query(ctx *gin.Context, model interface{}, db *gorm.DB, order string) (interface{}, Pagination, error) {
	pagination := GetPagination(ctx)
	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	err := db.Model(model).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order(order).Find(model).Error
	if err != nil {
		return nil, Pagination{}, err
	}
	return model, pagination, nil
}

type Pagination struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

// GetPagination 返回请求体中的分页信息
func GetPagination(ctx *gin.Context) Pagination {
	pagination := Pagination{}
	if err := ctx.BindJSON(&pagination); err != nil {
		pagination.PageNum = 1
		pagination.PageSize = 10
	}
	if pagination.PageNum == 0 {
		pagination.PageNum = 1
	}
	if pagination.PageSize > 100 {
		pagination.PageSize = 100
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	}
	return pagination
}
