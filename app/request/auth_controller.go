package request

import (
	"time"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32" required_msg:"用户名不能为空" min_msg:"用户名长度不能小于3" max_msg:"用户名长度不能大于32"`
	Password string `json:"password" binding:"required,min=5,max=32" required_msg:"密码不能为空" min_msg:"密码长度不能小于6" max_msg:"密码长度不能大于32"`
}
type ForgetPwdRequest struct {
	Email     string `json:"email" binding:"required,email" required_msg:"邮箱不能为空" email_msg:"邮箱格式不正确"`
	EmailCode string `json:"emailCode" binding:"required" required_msg:"邮箱验证码不能为空"`
	Password  string `json:"password" binding:"required,min=5,max=32" required_msg:"密码不能为空" min_msg:"密码长度不能小于6" max_msg:"密码长度不能大于32"`
}

type MessageRes struct {
	Content string `json:"content" binding:"required,min=3,max=255" required_msg:"留言内容不能为空" min_msg:"留言内容长度不能小于3" max_msg:"留言内容长度不能大于255"`
}

type Json struct {
	Username string    `json:"username"`
	CreatAt  time.Time `json:"creat_at"`
	UploadAt time.Time `json:"upload_at"`
}
