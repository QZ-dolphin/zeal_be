package admin

import (
	"context"

	v1 "zeal_be/api/admin/v1"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetSelfPrivs(ctx context.Context, req *v1.GetSelfPrivsReq) (res *v1.GetSelfPrivsRes, err error) {
	r := g.RequestFromCtx(ctx)
	email := r.GetCtxVar("Email").String()
	res = &v1.GetSelfPrivsRes{}
	res.Privs = service.Admin().GetUserPrivs(ctx, email)
	return
}
