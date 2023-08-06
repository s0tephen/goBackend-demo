package Image

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/middleware/auth"
	"net/http"
	"path/filepath"
	"time"
)

// UploadImage 上传图片
func UploadImage(ctx *gin.Context) {
	const maxUploadSize = 8 << 20 // 8MB
	err := ctx.Request.ParseMultipartForm(maxUploadSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("无法解析上传的表单", err.Error()))
		return
	}
	user := auth.CurrentUser(ctx)

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("无法获取图像", err.Error()))
		return
	}
	if file.Size > maxUploadSize {
		ctx.JSON(http.StatusBadRequest, response.New("文件太大", "文件大小不能超过8MB"))
		return
	}

	hash := md5.Sum([]byte(string(user.UID) + time.Now().String()))
	filename := hex.EncodeToString(hash[:]) + filepath.Ext(file.Filename)

	dest := filepath.Join("static/images", user.Username, "Image", filename)

	err = ctx.SaveUploadedFile(file, dest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("上传失败", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.New("上传成功", gin.H{"filename": filename}))
}
