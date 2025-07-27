package admin

import (
	"context"

	v1 "zeal_be/api/admin/v1"
	"zeal_be/internal/consts"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UpdateUserPrivs(ctx context.Context, req *v1.UpdateUserPrivsReq) (res *v1.UpdateUserPrivsRes, err error) {
	stat := service.Admin().CheckAdmin(ctx)
	if stat == consts.UserNoPrivilege {
		r := g.RequestFromCtx(ctx)
		r.Response.WriteStatusExit(404)
	}
	service.Admin().UpdateUserPrivs(ctx, req.T_email, req.Privs)
	res = &v1.UpdateUserPrivsRes{}
	res.Privs = service.Admin().GetUserPrivs(ctx, req.T_email)
	return
}
