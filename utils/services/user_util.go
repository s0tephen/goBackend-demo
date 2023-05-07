package services

import (
	"github.com/gin-gonic/gin"
	mysql "index_Demo/dao/mysql"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/utils/middleware/auth"
)

// IsAdmin 管理员权限验证
func IsAdmin(ctx *gin.Context) bool {
	user := auth.CurrentUser(ctx)
	u := dal.User
	userInfo, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	return userInfo.IsAdmin
}

type Pagination struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

// QueryPosts 获取分页信息
func QueryPosts(ctx *gin.Context) ([]model.Post, Pagination, error) {
	var posts []model.Post
	db := mysql.DB.GetDb()
	pagination := GetPagination(ctx)

	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	db.Model(&posts).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order("pTime desc").Find(&posts)
	if err := db.Error; err != nil {
		return nil, Pagination{}, nil
	}

	return posts, pagination, nil
}

// QueryUsers 返回用户列表和分页信息
func QueryUsers(ctx *gin.Context) ([]model.User, Pagination, error) {
	var users []model.User
	db := mysql.DB.GetDb()
	pagination := GetPagination(ctx)

	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	db.Model(&users).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order("create_at desc").Find(&users)
	if err := db.Error; err != nil {
		return nil, Pagination{}, nil
	}

	return users, pagination, nil
}

// QueryFdBack 返回反馈列表和分页信息
func QueryFdBack(ctx *gin.Context) ([]model.Feedback, Pagination, error) {
	var fdBack []model.Feedback
	db := mysql.DB.GetDb()
	pagination := GetPagination(ctx)

	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	db.Model(&fdBack).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order("create_at desc").Find(&fdBack)
	if err := db.Error; err != nil {
		return nil, Pagination{}, nil
	}

	return fdBack, pagination, nil
}

// GetPagination 返回请求体中的分页信息
func GetPagination(ctx *gin.Context) Pagination {
	pagination := Pagination{}
	if err := ctx.BindJSON(&pagination); err != nil {
		// 返回默认值
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
