package social

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/social/v1"
)

func (c *ControllerV1) GetSocialInfo(ctx context.Context, req *v1.GetSocialInfoReq) (res *v1.GetSocialInfoRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
