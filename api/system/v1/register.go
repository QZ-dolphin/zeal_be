package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta    `path:"/sys/register" method:"post"`
	NickName  string `p:"nickname" v:"required#昵称不能为空"`
	Email     string `p:"email" v:"required#邮箱不能为空"`
	Password  string `p:"password" v:"required#密码不能为空"`
	EmailCode string `p:"emailCode" v:"required#验证码不能为空"`
}

type RegisterRes struct {
	Stat bool `json:"stat"`
}
