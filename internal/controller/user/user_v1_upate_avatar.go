package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/user/v1"
)

func (c *ControllerV1) UpateAvatar(ctx context.Context, req *v1.UpateAvatarReq) (res *v1.UpateAvatarRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
