package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/user/v1"
)

func (c *ControllerV1) UserUpdatePwd(ctx context.Context, req *v1.UserUpdatePwdReq) (res *v1.UserUpdatePwdRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
