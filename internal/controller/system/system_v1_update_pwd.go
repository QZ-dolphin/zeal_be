package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/system/v1"
)

func (c *ControllerV1) UpdatePwd(ctx context.Context, req *v1.UpdatePwdReq) (res *v1.UpdatePwdRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
