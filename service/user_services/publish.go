package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"net/http"
	"strconv"
	"time"
)

// PublishPost 发布
func PublishPost(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	userPost := model.Post{}
	err := ctx.BindJSON(&userPost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("系统内部出错 请联系管理员", nil))
		return
	}

	post := model.Post{
		Username: &user.Username,
		PTitle:   userPost.PTitle,
		PCenter:  userPost.PCenter,
		PLabel:   userPost.PLabel,
		PTime:    time.Now(),
	}

	err = dal.Post.WithContext(ctx).Create(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("发布失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.New("发布成功", nil))

}

// ShowPost 展示文章
func ShowPost(ctx *gin.Context) {
	postID := ctx.Param("id")
	post, err := GetPostByID(postID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"post": post})
}

func GetPostByID(id string) (*model.Post, error) {
	// 根据文章ID从数据库或其他数据源查询文章信息的实现逻辑
	p := dal.Post

	postID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	post, err := p.Where(p.PID.Eq(int32(postID))).First()
	if err != nil {
		return nil, err
	}
	// 假设在这里执行查询操作，将结果赋值给 post 变量
	// 如果找不到文章，可以返回自定义的错误信息

	return post, nil
}
