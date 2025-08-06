package v1

import "github.com/gogf/gf/v2/frame/g"

type SetAnniversaryReq struct {
	g.Meta    `path:"/user/set_anniversary" method:"post"`
	Day_name  string `p:"day_name"`
	Day_value string `p:"day_value"`
}

type SetAnniversaryRes struct {
	Stat bool `json:"stat"`
}
