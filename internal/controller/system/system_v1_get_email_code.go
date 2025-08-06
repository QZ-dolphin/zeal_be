package system

import (
	"context"

	v1 "zeal_be/api/system/v1"
	"zeal_be/internal/service"
)

func (c *ControllerV1) GetEmailCode(ctx context.Context, req *v1.GetEmailCodeReq) (res *v1.GetEmailCodeRes, err error) {
	res = new(v1.GetEmailCodeRes)
	res.Send, res.TTL = service.System().EmailCodeSend2(ctx, req.ToEmail)
	return
}
