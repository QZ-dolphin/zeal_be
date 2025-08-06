package v1

import "github.com/gogf/gf/v2/frame/g"

type GetEmailCodeReq struct {
	g.Meta `path:"/user/get_email_code" method:"get"`
}

type GetEmailCodeRes struct {
	TTL  int64 `json:"ttl"` // 验证码有效时间，单位秒
	Send bool  `json:"send"`
}
