package system

import (
	"context"

	v1 "zeal_be/api/system/v1"
	"zeal_be/internal/consts"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UpdatePwd(ctx context.Context, req *v1.UpdatePwdReq) (res *v1.UpdatePwdRes, err error) {
	r := g.RequestFromCtx(ctx)
	verify_stat := service.System().EmailCodeVerify(ctx, req.Email, req.VerifyCode)
	if !verify_stat {
		r.Response.WriteJsonExit(g.Map{"code": 51, "msg": "验证码错误"})
	}
	exist_stat := service.System().UpdatePwd(ctx, req.Email, req.NewPwd)
	if exist_stat == consts.UserNotExist {
		r.Response.WriteJsonExit(g.Map{"code": 51, "msg": "用户不存在"})
	}
	res = &v1.UpdatePwdRes{}
	res.Stat = true
	return
}
