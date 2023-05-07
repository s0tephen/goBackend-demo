package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"index_Demo/gen/orm/dal"
	"index_Demo/utils/middleware/auth"
)

// IsAdmin 管理员权限验证
func IsAdmin(ctx *gin.Context) bool {
	user := auth.CurrentUser(ctx)
	u := dal.User
	userInfo, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	return userInfo.IsAdmin
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
