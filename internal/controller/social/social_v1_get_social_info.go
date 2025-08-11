package social

import (
	"context"

	v1 "zeal_be/api/social/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) GetSocialInfo(ctx context.Context, req *v1.GetSocialInfoReq) (res *v1.GetSocialInfoRes, err error) {
	if !service.Social().IsInSocial(ctx) {
		service.Social().CreateSocialAccount(ctx)
	}
	res = &v1.GetSocialInfoRes{}
	// res.Doc_count = service.Notebooks().GetBookCount(ctx)
	res.Following, res.Followers = service.Social().GetSocialInfo(ctx)
	return
}
