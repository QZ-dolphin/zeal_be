package admin

import (
	"context"

	v1 "zeal_be/api/admin/v1"
	"zeal_be/internal/consts"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetUserPrivs(ctx context.Context, req *v1.GetUserPrivsReq) (res *v1.GetUserPrivsRes, err error) {
	stat := service.Admin().CheckAdmin(ctx)
	if stat == consts.UserNoPrivilege {
		r := g.RequestFromCtx(ctx)
		r.Response.WriteStatusExit(404)
	}
	res = &v1.GetUserPrivsRes{}
	res.Privs = service.Admin().GetUserPrivs(ctx, req.T_email)
	return
}
