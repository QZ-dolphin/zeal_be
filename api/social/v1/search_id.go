package v1

import "github.com/gogf/gf/v2/frame/g"

type SearchIdReq struct {
	g.Meta `path:"/social/search_id" method:"post"`
	SearchReqInfo
}

type SearchReqInfo struct {
	ID_desc   string `p:"id_desc"`
	PageSize  int    `p:"pageSize "`
	Page      int    `p:"pageIndex"`
	NeedCount int    `p:"needCount"`
}

type SearchIdRes struct {
	User_list []User_info `json:"user_list"`
	Total     int         `json:"total"`
}

type User_info struct {
	User_id       int    `json:"user_id"`
	User_name     string `json:"user_name"`
	Email_address string `json:"email_address"`
	Avatar_url    string `json:"avatar_url"`
	IsFollowing   bool   `json:"isFollowing"`
	IsFollower    bool   `json:"isFollower"`
}
