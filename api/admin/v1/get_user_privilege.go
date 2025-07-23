package v1

import "github.com/gogf/gf/v2/frame/g"

type GetUserPrivsReq struct {
	g.Meta  `path:"/admin/get_user_privs" method:"post"`
	T_email string `p:"target_email"`
}

type GetUserPrivsRes struct {
	Privs []string `json:"privs_list"`
}
