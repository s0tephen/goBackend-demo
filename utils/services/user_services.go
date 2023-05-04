package services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"index_Demo/app/request"
	"index_Demo/dao/redisServer"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"time"
)

func UserExist(ctx *gin.Context, username string) bool {
	u := dal.User
	user, _ := u.WithContext(ctx).Where(u.Username.Eq(username)).First()
	return user != nil
}

func GetCodeFromRedis(email string) (string, error) {
	return redisServer.Get(email)
}

func EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CreateUser(regRequest request.RegRequest, avatar string, hashPassword string) *model.User {
	createTime := time.Now()
	return &model.User{
		Username: regRequest.Username,
		Avatar:   &avatar,
		Uemail:   &regRequest.Email,
		Password: hashPassword,
		CreateAt: createTime,
	}
}

func SaveUser(user *model.User) error {
	return dal.User.Create(user)
}
