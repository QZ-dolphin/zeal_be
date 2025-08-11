package v1

import "github.com/gogf/gf/v2/frame/g"

type Userdata struct {
	Username    string `bson:"user_name"`
	UserId      int    `bson:"user_id"`
	Nickname    string `bson:"nick_name"`
	Avatar_url  string `bson:"avatar_url"`
	IsFollowing bool   `json:"isFollowing"`
	IsFollower  bool   `json:"isFollower"`
}

type GetFollowerListReq struct {
	g.Meta    `path:"/social/get_follower_list" method:"post"`
	PageSize  int  `p:"pageSize "`
	Page      int  `p:"pageIndex"`
	NeedCount bool `p:"needCount"`
}

type GetFollowerListRes struct {
	UserList []Userdata `json:"user_list"`
	Total    int        `json:"total"`
}
