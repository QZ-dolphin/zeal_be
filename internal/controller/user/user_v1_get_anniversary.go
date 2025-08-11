package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) GetAnniversary(ctx context.Context, req *v1.GetAnniversaryReq) (res *v1.GetAnniversaryRes, err error) {
	res = &v1.GetAnniversaryRes{}
	res.Day_name, res.Day_value = service.User().GetAnniversary(ctx)
	return
}
