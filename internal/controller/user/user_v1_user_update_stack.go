package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) UserUpdateStack(ctx context.Context, req *v1.UserUpdateStackReq) (res *v1.UserUpdateStackRes, err error) {
	res = &v1.UserUpdateStackRes{}
	res.Stat, res.Update_at = service.User().UpdateStackText(ctx, req.Stack_text)
	return
}
