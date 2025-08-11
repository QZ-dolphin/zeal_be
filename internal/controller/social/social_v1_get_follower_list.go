package social

import (
	"context"

	v1 "zeal_be/api/social/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) GetFollowerList(ctx context.Context, req *v1.GetFollowerListReq) (res *v1.GetFollowerListRes, err error) {
	res = &v1.GetFollowerListRes{}
	res.UserList = service.Social().GetFollowerList(ctx, req.PageSize, req.Page)
	if req.NeedCount {
		_, res.Total = service.Social().GetSocialInfo(ctx)
	}
	return
}
