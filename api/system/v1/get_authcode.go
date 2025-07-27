package v1

import "github.com/gogf/gf/v2/frame/g"

type GetAuthCodeReq struct {
	g.Meta `path:"/sys/get_authcode" method:"get"`
}
type GetAuthCodeRes struct {
	Key string `json:"key"`
	Img string `json:"img"`
}
