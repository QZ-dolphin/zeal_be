package social

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/social/v1"
)

func (c *ControllerV1) SearchId(ctx context.Context, req *v1.SearchIdReq) (res *v1.SearchIdRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
