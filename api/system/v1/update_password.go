package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdatePwdReq struct {
	g.Meta     `path:"/sys/update_pwd" method:"post"`
	Email      string `p:"email"`
	VerifyCode string `p:"verify_code"`
	NewPwd     string `p:"new_pwd"`
}

type UpdatePwdRes struct {
	Stat bool `json:"stat"`
}
