package v1

import "github.com/gogf/gf/v2/frame/g"

type UserUpdateStackReq struct {
	g.Meta     `path:"/user/update_stack" method:"post"`
	Stack_text string `p:"stack_text"`
}

type UserUpdateStackRes struct {
	Stat      bool   `json:"stat"`
	Update_at string `json:"update_at"`
}
