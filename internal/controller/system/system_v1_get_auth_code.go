package system

import (
	"context"

	v1 "zeal_be/api/system/v1"
	"zeal_be/internal/service"
	"zeal_be/utility"
)

func (c *ControllerV1) GetAuthCode(ctx context.Context, req *v1.GetAuthCodeReq) (res *v1.GetAuthCodeRes, err error) {
	idKeyC, base64stringC, err := service.System().GetVerifyImgString(ctx)
	if err != nil {
		utility.Clog("\n获取验证码失败", err)
		return nil, err
	}
	res = new(v1.GetAuthCodeRes)
	res.Key = idKeyC
	res.Img = base64stringC
	return
}
