package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/admin/v1"
)

func (c *ControllerV1) UpdateUserPrivs(ctx context.Context, req *v1.UpdateUserPrivsReq) (res *v1.UpdateUserPrivsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
