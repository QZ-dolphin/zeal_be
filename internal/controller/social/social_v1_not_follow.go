package social

import (
	"context"

	v1 "zeal_be/api/social/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) NotFollow(ctx context.Context, req *v1.NotFollowReq) (res *v1.NotFollowRes, err error) {
	res = &v1.NotFollowRes{}
	res.Stat = service.Social().NotFollowing(ctx, req.User_id)
	return
}
