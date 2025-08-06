package system

import (
	"context"

	v1 "zeal_be/api/system/v1"
	"zeal_be/internal/consts"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	verify_stat := service.System().EmailCodeVerify(ctx, req.Email, req.EmailCode)
	r := g.RequestFromCtx(ctx)
	if !verify_stat {
		r.Response.WriteJsonExit(g.Map{"code": 51, "msg": "验证码错误"})
	}
	register_stat := service.System().Register(ctx, req.NickName, req.Email, req.Password)
	if register_stat == consts.UserAlreadyExist {
		r.Response.WriteJsonExit(g.Map{"code": 51, "msg": "邮箱已注册"})
	}
	res = new(v1.RegisterRes)
	res.Stat = true
	return
}
