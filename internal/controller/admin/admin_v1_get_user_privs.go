package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/admin/v1"
)

func (c *ControllerV1) GetUserPrivs(ctx context.Context, req *v1.GetUserPrivsReq) (res *v1.GetUserPrivsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
