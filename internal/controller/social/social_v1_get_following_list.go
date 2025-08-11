package social

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/social/v1"
)

func (c *ControllerV1) GetFollowingList(ctx context.Context, req *v1.GetFollowingListReq) (res *v1.GetFollowingListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
