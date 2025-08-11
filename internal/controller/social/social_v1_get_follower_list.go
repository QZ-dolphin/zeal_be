package social

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/social/v1"
)

func (c *ControllerV1) GetFollowerList(ctx context.Context, req *v1.GetFollowerListReq) (res *v1.GetFollowerListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
