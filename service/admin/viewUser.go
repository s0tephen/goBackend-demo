package admin

import (
	"github.com/gin-gonic/gin"
	mysql "goBackend-demo/dao/mysql"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/services"
	"net/http"
)

type EditUserInfo struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func ViewUserList(ctx *gin.Context) {
	if !services.IsAdmin(ctx) {
		ctx.JSON(http.StatusUnauthorized, response.New("未授权", nil))
		return
	}
	// 根据请求参数判断是查看用户列表还是修改用户信息
	if ctx.Query("action") == "edit" {
		var userInfo EditUserInfo
		err := ctx.BindJSON(&userInfo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.New(err.Error(), nil))
			return
		}
		// 调用修改用户信息的服务
		editUser, err := services.EditUser(ctx, userInfo.UserID, userInfo.UserName, userInfo.Password)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
			return
		}
		ctx.JSON(http.StatusOK, response.New("修改成功", editUser))
	} else {
		queryUsers, pagination, err := services.Query(ctx, &[]model.User{}, mysql.DB.GetDb(), "create_at")
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
			return
		}
		ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
			"list":     queryUsers,
			"total":    pagination.Total,
			"pages":    pagination.Total / int64(pagination.PageSize),
			"pageNum":  pagination.PageNum,
			"pageSize": pagination.PageSize,
		}))
	}
}

func DeleteUser(ctx *gin.Context) {
	if !services.IsAdmin(ctx) {
		ctx.JSON(http.StatusUnauthorized, response.New("未授权", nil))
		return
	}
	userID := ctx.Query("Uid")
	info, err := services.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("删除成功", info))
}

// ViewUserLists 查看用户列表
func ViewUserLists(ctx *gin.Context) {
	if !services.IsAdmin(ctx) {
		ctx.JSON(http.StatusUnauthorized, response.New("Unauthorized", nil))
		return
	}
	queryUsers, pagination, err := services.Query(ctx, &[]model.User{}, mysql.DB.GetDb(), "create_at")
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
		"list":     queryUsers,
		"total":    pagination.Total,
		"pages":    pagination.Total / int64(pagination.PageSize),
		"pageNum":  pagination.PageNum,
		"pageSize": pagination.PageSize,
	}))
}
