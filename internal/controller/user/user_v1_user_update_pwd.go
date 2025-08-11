package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UserUpdatePwd(ctx context.Context, req *v1.UserUpdatePwdReq) (res *v1.UserUpdatePwdRes, err error) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	verify_stat := service.System().EmailCodeVerify(ctx, user_email, req.Verify_code)
	if !verify_stat {
		r.Response.WriteJsonExit(g.Map{"code": 51, "msg": "验证码错误"})
	}
	service.System().UpdatePwd(ctx, user_email, req.NewPwd)
	res = &v1.UserUpdatePwdRes{}
	res.Stat = true
	return
}
