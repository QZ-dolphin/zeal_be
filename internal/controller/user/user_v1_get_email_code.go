package user

import (
	"context"

	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetEmailCode(ctx context.Context, req *v1.GetEmailCodeReq) (res *v1.GetEmailCodeRes, err error) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	res = &v1.GetEmailCodeRes{}
	res.Send, res.TTL = service.System().EmailCodeSend2(ctx, user_email)
	return
}
