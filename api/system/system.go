// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"zeal_be/api/system/v1"
)

type ISystemV1 interface {
	GetAuthCode(ctx context.Context, req *v1.GetAuthCodeReq) (res *v1.GetAuthCodeRes, err error)
	GetEmailCode(ctx context.Context, req *v1.GetEmailCodeReq) (res *v1.GetEmailCodeRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	UpdatePwd(ctx context.Context, req *v1.UpdatePwdReq) (res *v1.UpdatePwdRes, err error)
}
