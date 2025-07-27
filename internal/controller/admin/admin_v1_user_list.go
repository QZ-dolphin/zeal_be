package admin

import (
	"context"

	v1 "zeal_be/api/admin/v1"
	"zeal_be/internal/consts"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	var stat int
	res = &v1.UserListRes{}
	stat, res.User_list, res.Count_num = service.Admin().GetUserList(ctx, req.ULReqInfo)
	if stat == consts.UserNoPrivilege {
		r := g.RequestFromCtx(ctx)
		r.Response.WriteStatusExit(404)
	}
	return
}
