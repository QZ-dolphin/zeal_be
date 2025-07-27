package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/sys/login" method:"post"`
	Email    string `p:"email" v:"required#邮箱不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
	// VerifyKey  string `p:"verifyKey" v:"required#验证码不能为空"`
	// VerifyCode string `p:"verifyCode"`
}

type LoginRes struct {
	Token string `json:"token"`
}
