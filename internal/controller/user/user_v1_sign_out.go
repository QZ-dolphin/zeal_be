package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	res = &v1.SignOutRes{}
	return
}
