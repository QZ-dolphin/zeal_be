package v1

import "github.com/gogf/gf/v2/frame/g"

type GetSelfPrivsReq struct {
	g.Meta `path:"/admin/get_self_privs" method:"get"`
}

type GetSelfPrivsRes struct {
	Privs []string `json:"privs_list"`
}
