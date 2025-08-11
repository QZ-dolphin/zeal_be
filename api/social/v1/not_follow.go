package v1

import "github.com/gogf/gf/v2/frame/g"

type NotFollowReq struct {
	g.Meta  `path:"/social/not_follow" method:"post"`
	User_id int `json:"user_id"`
}

type NotFollowRes struct {
	Stat bool `json:"stat"`
}
