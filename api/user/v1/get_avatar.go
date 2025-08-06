package v1

import "github.com/gogf/gf/v2/frame/g"

type GetAvatarReq struct {
	g.Meta `path:"/user/get_avatar" method:"get"`
}

type GetAvatarRes struct {
	Filepath string `json:"avatar_url"`
}
