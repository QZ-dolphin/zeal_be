package v1

import "github.com/gogf/gf/v2/frame/g"

type GetEmailCodeReq struct {
	g.Meta  `path:"/sys/get_emailcode" method:"post"`
	ToEmail string `json:"toemail" v:"email#Email格式错误"`
}

type GetEmailCodeRes struct {
	TTL  int64 `json:"ttl"` // 验证码有效时间，单位秒
	Send bool  `json:"send"`
}
