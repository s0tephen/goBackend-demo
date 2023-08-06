package services

import (
	"github.com/gin-gonic/gin"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/middleware/auth"
	"gorm.io/gorm"
	"strconv"
)

// IsAdmin Admin管理员权限验证
func IsAdmin(ctx *gin.Context) bool {
	user := auth.CurrentUser(ctx)
	u := dal.User
	userInfo, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	return userInfo.IsAdmin
}

// DeleteUser Admin删除用户
func DeleteUser(ctx *gin.Context, userID string) (interface{}, error) {
	u := dal.User
	uID, _ := strconv.Atoi(userID)
	DeleteInfo, err := u.WithContext(ctx).Where(u.UID.Eq(int32(uID))).Delete()
	if err != nil {
		return response.New("删除失败", err.Error()), nil
	}
	if DeleteInfo.Error != nil && DeleteInfo.RowsAffected == 0 {
		return response.New("删除失败", DeleteInfo.Error.Error()), nil
	}
	return response.New("删除成功", DeleteInfo.RowsAffected), nil
}

// EditUser Admin修改用户信息
func EditUser(ctx *gin.Context, userID string, userName string, password string) (*model.User, error) {
	u := dal.User
	uID, _ := strconv.Atoi(userID)
	Info, err := u.WithContext(ctx).Where(u.UID.Eq(int32(uID))).First()
	if err != nil {
		return &model.User{}, err
	}
	Info.Username = userName
	pwd, _ := EncryptPassword(password)
	Info.Password = pwd
	_, err = u.WithContext(ctx).Where(u.UID.Eq(int32(uID))).Updates(Info)
	if err != nil {
		return &model.User{}, err
	}
	return Info, nil
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
