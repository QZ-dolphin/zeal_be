package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) UpateAvatar(ctx context.Context, req *v1.UpateAvatarReq) (res *v1.UpateAvatarRes, err error) {
	res = &v1.UpateAvatarRes{}
	res.Avatar_url = service.User().UpateAvatar(ctx, req.File)
	return
}
