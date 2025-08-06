package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/user/v1"
)

func (c *ControllerV1) GetAvatar(ctx context.Context, req *v1.GetAvatarReq) (res *v1.GetAvatarRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
