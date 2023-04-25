package admin_services

import (
	"github.com/gin-gonic/gin"
	mysql "index_Demo/dao/mysql"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/userUtil"
	"net/http"
)

type Pagination struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

func ViewUserList(ctx *gin.Context) {
	var userSql []model.User
	//u := dal.User
	db := mysql.DB.GetDb()

	//查询数据库该用户是否为管理员
	if userUtil.IsAdmin(ctx) == true {
		ctx.JSON(http.StatusUnauthorized, response.New("Unauthorized", nil))
		return
	}

	queryUsers, pagination, err := userUtil.QueryUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
		return
	}

	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	db.Model(userSql).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order("create_at desc").Find(&userSql)
	ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
		"list":     queryUsers,
		"total":    pagination.Total,
		"pages":    pagination.Total / int64(pagination.PageSize),
		"pageNum":  pagination.PageNum,
		"pageSize": pagination.PageSize,
	}))
}
