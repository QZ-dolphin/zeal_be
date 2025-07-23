package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/admin/v1"
)

func (c *ControllerV1) GetSelfPrivs(ctx context.Context, req *v1.GetSelfPrivsReq) (res *v1.GetSelfPrivsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
