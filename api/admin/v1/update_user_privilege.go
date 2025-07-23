package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateUserPrivsReq struct {
	g.Meta  `path:"/admin/update_user_privs" method:"post"`
	T_email string   `p:"target_email"`
	Privs   []string `p:"privs_list"`
}

type UpdateUserPrivsRes struct {
	Privs []string `json:"privs_list"`
}
