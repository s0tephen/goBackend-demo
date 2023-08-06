package user

import (
	"github.com/gin-gonic/gin"
	mysql "goBackend-demo/dao/mysql"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/middleware/auth"
	"goBackend-demo/utils/services"
	"net/http"
	"strconv"
	"time"
)

func PostList(ctx *gin.Context) {
	queryPosts, pagination, err := services.Query(ctx, &[]model.Post{}, mysql.DB.GetDb(), "pTime")
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
		"list":     queryPosts,
		"total":    pagination.Total,
		"pages":    pagination.Total / int64(pagination.PageSize),
		"pageNum":  pagination.PageNum,
		"pageSize": pagination.PageSize,
	}))
}

// PostDetail 文章内容
func PostDetail(ctx *gin.Context) {
	postID := ctx.Query("id")
	post, err := GetPostByID(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"post": post})
}

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

func GetPostByID(ctx *gin.Context, postId string) (*model.Post, error) {
	p := dal.Post

	id, err := strconv.Atoi(postId)
	if err != nil {
		return nil, err
	}

	post, err := p.WithContext(ctx).Where(p.PID.Eq(int32(id))).First()
	if err != nil {
		return nil, err
	}
	return post, nil
}
