package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/app/request"
	"index_Demo/gen/response"
	"net/http"
)

func RegEmailCode(ctx *gin.Context) {
	regEmail := request.RegRequest{}
	err := ctx.BindJSON(&regEmail)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("系统出错联系管理员", err))
		return
	}
}
