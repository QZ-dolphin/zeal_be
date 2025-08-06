package v1

import "github.com/gogf/gf/v2/frame/g"

type GetSelfInfoReq struct {
	g.Meta `path:"/user/get_self_info" method:"get"`
}

type GetSelfInfoRes struct {
	Info SelfInfo `json:"self_info"`
}

type SelfInfo struct {
	User_name       string `json:"user_name"`
	Email_address   string `json:"email_address"`
	Created_time    string `json:"created_time"`
	Last_login_time string `json:"last_login_time"`
	Avatar_url      string `json:"avatar_url"`
}
