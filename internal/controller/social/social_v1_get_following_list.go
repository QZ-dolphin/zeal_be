package social

import (
	"context"

	v1 "zeal_be/api/social/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) GetFollowingList(ctx context.Context, req *v1.GetFollowingListReq) (res *v1.GetFollowingListRes, err error) {
	res = &v1.GetFollowingListRes{}
	res.UserList = service.Social().GetFollowingList(ctx, req.PageSize, req.Page)
	if req.NeedCount {
		res.Total, _ = service.Social().GetSocialInfo(ctx)
	}
	return
}
