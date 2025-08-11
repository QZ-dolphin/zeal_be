package social

import (
	"context"

	v1 "zeal_be/api/social/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) Follow(ctx context.Context, req *v1.FollowReq) (res *v1.FollowRes, err error) {
	res = &v1.FollowRes{}
	res.Stat = service.Social().ToFollowing(ctx, req.User_id, req.User_name, req.Nick_name, req.Avatar_url)
	return
}
