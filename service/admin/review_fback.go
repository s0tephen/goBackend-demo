package admin

import (
	"github.com/gin-gonic/gin"
	mysql "index_Demo/dao/mysql"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/services"
	"net/http"
)

// ReviewFeedback 查询反馈
func ReviewFeedback(ctx *gin.Context) {
	if !services.IsAdmin(ctx) {
		ctx.JSON(http.StatusUnauthorized, response.New("Unauthorized", nil))
		return
	}
	QueryFdBack, pagination, err := services.Query(ctx, &[]model.Feedback{}, mysql.DB.GetDb(), "fTime")
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
		"list":     QueryFdBack,
		"total":    pagination.Total,
		"pages":    pagination.Total / int64(pagination.PageSize),
		"pageNum":  pagination.PageNum,
		"pageSize": pagination.PageSize,
	}))

}
