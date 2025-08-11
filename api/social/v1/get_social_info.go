package v1

import "github.com/gogf/gf/v2/frame/g"

type GetSocialInfoReq struct {
	g.Meta `path:"/social/get_social_info" method:"get"`
}

type Social_info struct {
	Doc_count int `json:"doc_count"`
	Following int `json:"following"`
	Followers int `json:"followers"`
}

type GetSocialInfoRes struct {
	Social_info
}
