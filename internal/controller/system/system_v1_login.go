package system

import (
	"context"

	v1 "zeal_be/api/system/v1"
	"zeal_be/internal/consts"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	r := g.RequestFromCtx(ctx)
	res = new(v1.LoginRes)
	stat, token := service.System().Login(ctx, req.Email, req.Password)
	if stat == consts.UserNotExist {
		r.Response.WriteJsonExit(g.Map{"code": 51, "msg": "用户不存在"})
	} else if stat == consts.UserFalsePassword {
		r.Response.WriteJsonExit(g.Map{"code": 51, "msg": "密码错误"})
	}
	res.Token = token
	return
}
