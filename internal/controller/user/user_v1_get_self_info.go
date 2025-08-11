package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) GetSelfInfo(ctx context.Context, req *v1.GetSelfInfoReq) (res *v1.GetSelfInfoRes, err error) {
	res = &v1.GetSelfInfoRes{}
	res.Info = *service.User().GetSelfInfo(ctx)
	return
}
