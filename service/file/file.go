package file

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"goBackend-demo/gen/orm/dal"
	"goBackend-demo/gen/orm/model"
	"goBackend-demo/gen/response"
	"goBackend-demo/utils/middleware/auth"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func UploadFile(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("Failed to save avatar file", err.Error()))
		return
	}
	files := form.File["files"]
	tx := dal.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var f multipart.File
	var data []byte
	filePaths := make([]string, 0)

	// 获取请求的协议、域名和端口
	request := ctx.Request
	host := request.Host
	scheme := "http"
	if request.TLS != nil {
		scheme = "https"
	}

	for _, file := range files {
		f, err = file.Open()
		if err != nil {
			break
		}
		data, err = io.ReadAll(f)
		if err != nil {
			break
		}
		md5Str := sumMd5(data)
		fileM := &model.File{}
		_, err = tx.WithContext(ctx).File.Where(tx.File.Md5.Eq(md5Str)).First()

		// 构建文件保存的目标路径，相对于static/data
		fileExt := filepath.Ext(file.Filename)
		filePath := path.Join("static/data", md5Str+fileExt)

		if err == nil {
			// 移动文件到目标路径
			err = os.WriteFile(filePath, data, os.ModePerm)
			if err != nil {
				break
			}
		}

		fileM.Md5 = md5Str
		fileM.Uploaded = user.Username
		fileM.Filename = file.Filename
		fileM.Size = int32(file.Size)
		fileM.UpTime = time.Now()
		err = tx.File.Create(fileM)
		if err != nil {
			break
		}

		// 构建完整的URL
		fileURL := fmt.Sprintf("%s://%s/%s", scheme, host, filePath)
		filePaths = append(filePaths, fileURL)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("err:", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": filePaths,
	})
}

func sumMd5(data []byte) string {
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum)
}
