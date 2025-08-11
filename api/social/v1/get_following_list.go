package v1

import "github.com/gogf/gf/v2/frame/g"

type GetFollowingListReq struct {
	g.Meta    `path:"/social/get_following_list" method:"post"`
	PageSize  int  `p:"pageSize "`
	Page      int  `p:"pageIndex"`
	NeedCount bool `p:"needCount"`
}

type GetFollowingListRes struct {
	UserList []Userdata `json:"user_list"`
	Total    int        `json:"total"`
}
