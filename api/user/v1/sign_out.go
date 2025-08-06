package v1

import "github.com/gogf/gf/v2/frame/g"

type SignOutReq struct {
	g.Meta `path:"/user/sign_out" method:"get"`
}

type SignOutRes struct{}
