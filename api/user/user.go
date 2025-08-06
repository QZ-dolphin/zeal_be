// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"zeal_be/api/user/v1"
)

type IUserV1 interface {
	GetAnniversary(ctx context.Context, req *v1.GetAnniversaryReq) (res *v1.GetAnniversaryRes, err error)
	GetAvatar(ctx context.Context, req *v1.GetAvatarReq) (res *v1.GetAvatarRes, err error)
	GetEmailCode(ctx context.Context, req *v1.GetEmailCodeReq) (res *v1.GetEmailCodeRes, err error)
	GetSelfInfo(ctx context.Context, req *v1.GetSelfInfoReq) (res *v1.GetSelfInfoRes, err error)
	UserGetStack(ctx context.Context, req *v1.UserGetStackReq) (res *v1.UserGetStackRes, err error)
	SetAnniversary(ctx context.Context, req *v1.SetAnniversaryReq) (res *v1.SetAnniversaryRes, err error)
	SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error)
	UpateAvatar(ctx context.Context, req *v1.UpateAvatarReq) (res *v1.UpateAvatarRes, err error)
	UserUpdatePwd(ctx context.Context, req *v1.UserUpdatePwdReq) (res *v1.UserUpdatePwdRes, err error)
	UserUpdateStack(ctx context.Context, req *v1.UserUpdateStackReq) (res *v1.UserUpdateStackRes, err error)
}
