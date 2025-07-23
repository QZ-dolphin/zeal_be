package v1

import "github.com/gogf/gf/v2/frame/g"

type UserListReq struct {
	g.Meta `path:"/admin/get_user_list" method:"post"`
	ULReqInfo
}

type ULReqInfo struct {
	PageSize  int    `p:"pageSize "`
	PageIndex int    `p:"pageIndex"`
	Order     string `p:"order"`
	OrderItem string `p:"orderItem"`
	NeedCount int    `p:"needCount"`
}

type UserInfo struct {
	User_ID         int    `json:"user_ID"`
	User_name       string `json:"user_name"`
	Email_address   string `json:"email_address"`
	Password        string `json:"password"`
	Created_time    string `json:"created_time"`
	Last_login_time string `json:"last_login_time"`
	Avatar_url      string `json:"avatar_url"`
	Ip_address      string `json:"ip_address"`
}

type UserListRes struct {
	User_list []UserInfo `json:"user_list"`
	Count_num int        `json:"count_num"`
}
