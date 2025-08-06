package v1

import "github.com/gogf/gf/v2/frame/g"

type UserUpdatePwdReq struct {
	g.Meta      `path:"/user_update_pwd" method:"post"`
	Verify_code string `p:"verify_code"`
	NewPwd      string `p:"new_pwd"`
}

type UserUpdatePwdRes struct {
	Stat bool `json:"stat"`
}
