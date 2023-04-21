package user_services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"index_Demo/app/request"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/validateUtils"
	"net/http"
	"time"
)

func Register(ctx *gin.Context) {

	regRequest := request.RegRequest{}
	err := ctx.BindJSON(&regRequest)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, response.New("系统内部出错 请联系管理员", err.Error()))
	//	return
	//}
	message, hasErr := validateUtils.ReturnValidateMessage(&regRequest, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}

	u := dal.User
	user, _ := u.WithContext(ctx).Where(u.Username.Eq(regRequest.Username)).First()
	if user != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("用户已存在", nil))
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(regRequest.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("密码加密失败", err.Error()))
		return
	}

	user = &model.User{
		Username: regRequest.Username,
		Password: string(hashPassword),
		CreateAt: time.Now(),
	}

	err = dal.User.Create(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("注册失败", err.Error()))
	}

	type UserJson struct {
		Username string    `json:"username"`
		CreatAt  time.Time `json:"creat_at"`
	}
	userJson := UserJson{
		Username: regRequest.Username,
		CreatAt:  time.Now(),
	}
	ctx.JSON(http.StatusOK, response.New("注册成功", userJson))
}
