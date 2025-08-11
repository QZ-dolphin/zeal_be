package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) GetAvatar(ctx context.Context, req *v1.GetAvatarReq) (res *v1.GetAvatarRes, err error) {
	res = &v1.GetAvatarRes{}
	res.Filepath = service.User().GetAvatar(ctx)
	return
}
