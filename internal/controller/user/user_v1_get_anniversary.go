package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"zeal_be/api/user/v1"
)

func (c *ControllerV1) GetAnniversary(ctx context.Context, req *v1.GetAnniversaryReq) (res *v1.GetAnniversaryRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
