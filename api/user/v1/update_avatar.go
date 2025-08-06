package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UpateAvatarReq struct {
	g.Meta `path:"/user/upload_avatar" method:"post" mime:"multipart/form-data"`
	File   *ghttp.UploadFile `p:"ufile" type:"file"`
}

type UpateAvatarRes struct {
	Avatar_url string `json:"avatar_url"`
}
