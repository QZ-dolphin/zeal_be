package v1

import "github.com/gogf/gf/v2/frame/g"

type FollowReq struct {
	g.Meta     `path:"/social/follow" method:"post"`
	User_id    int    `json:"user_id"`
	User_name  string `json:"user_name"`
	Nick_name  string `json:"nick_name"`
	Avatar_url string `json:"avatar_url"`
}

type FollowRes struct {
	Stat bool `json:"stat"`
}
