// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "zeal_be/api/user/v1"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IUser interface {
		UpateAvatar(ctx context.Context, file *ghttp.UploadFile) (filepath string)
		GetAvatar(ctx context.Context) (avatar string)
		SetAnniversary(ctx context.Context, day_name string, day_value string)
		GetAnniversary(ctx context.Context) (day_name string, day_value string)
		GetSelfInfo(ctx context.Context) (infoData *v1.SelfInfo)
		UpdateStackText(ctx context.Context, stacktext string) (stat bool, update_at string)
		GetStackText(ctx context.Context) (stacktext string, update_at string)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
