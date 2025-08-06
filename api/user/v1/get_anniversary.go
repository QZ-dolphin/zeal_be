package v1

import "github.com/gogf/gf/v2/frame/g"

type GetAnniversaryReq struct {
	g.Meta `path:"/user/get_anniversary" method:"get"`
}

type GetAnniversaryRes struct {
	Day_name  string `json:"day_name"`
	Day_value string `json:"day_value"`
}
