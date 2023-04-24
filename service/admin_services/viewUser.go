package admin_services

import (
	"github.com/gin-gonic/gin"
	mysql "index_Demo/dao/mysql"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"net/http"
)

type Pagination struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

func ViewUser(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	var userSql model.User
	u := dal.User
	db := mysql.DB.GetDb()

	//查询数据库该用户是否为管理员
	userInfo, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	if userInfo.IsAdmin == false {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("Unauthorized", nil))
		return
	}
	pagination := Pagination{}
	err := ctx.BindJSON(&pagination)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("绑定数据失败", nil))
		return
	}

	if pagination.PageNum == 0 {
		pagination.PageNum = 1
	}
	switch {
	case pagination.PageSize > 100:
		pagination.PageSize = 100
	case pagination.PageSize <= 0:
		pagination.PageSize = 10
	}
	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	db.Model(userSql).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order("create_at desc").Find(&userSql)
	ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
		"list":     userSql,
		"total":    pagination.Total,
		"pageNum":  pagination.PageNum,
		"pageSize": pagination.PageSize,
		"pages":    pagination.Total / int64(pagination.PageSize),
	}))
}
