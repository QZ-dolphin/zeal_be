package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) SetAnniversary(ctx context.Context, req *v1.SetAnniversaryReq) (res *v1.SetAnniversaryRes, err error) {
	service.User().SetAnniversary(ctx, req.Day_name, req.Day_value)
	return
}
