package v1

import "github.com/gogf/gf/v2/frame/g"

type UserGetStackReq struct {
	g.Meta `path:"/user/get_stack" method:"get"`
}

type UserGetStackRes struct {
	Stack_text string `json:"stack_text"`
	Update_at  string `json:"update_at"`
}
