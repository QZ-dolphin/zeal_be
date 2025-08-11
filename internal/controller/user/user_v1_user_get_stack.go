package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) UserGetStack(ctx context.Context, req *v1.UserGetStackReq) (res *v1.UserGetStackRes, err error) {
	res = &v1.UserGetStackRes{}
	res.Stack_text, res.Update_at = service.User().GetStackText(ctx)
	return
}
