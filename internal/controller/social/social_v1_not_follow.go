package social

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/social/v1"
)

func (c *ControllerV1) NotFollow(ctx context.Context, req *v1.NotFollowReq) (res *v1.NotFollowRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
