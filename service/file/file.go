package file

import (
	"crypto/md5"
	"fmt"
	logs "github.com/danbai225/go-logs"
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
	filePath := make([]string, 0)
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
		// 返回URL 为了方便前端直接使用
		// Get the file extension from the uploaded file
		fileExt := filepath.Ext(file.Filename)
		logs.InfoF("fileExt:%s", fileExt)
		fileM.Path = fmt.Sprintf("http://localhost:9990/static/data/%s%s", md5Str, fileExt)
		//fileM.Path = fmt.Sprintf("static/data%c%s", os.PathSeparator, md5Str)
		makeAll(fileM.Path)
		if err == nil {
			err = os.WriteFile(fileM.Path, data, os.ModePerm)
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
		filePath = append(filePath, fileM.Path)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("err:", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": filePath,
	})
}
func sumMd5(data []byte) string {
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum)
}
func makeAll(p string) {
	err := os.MkdirAll(path.Dir(p), 644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
