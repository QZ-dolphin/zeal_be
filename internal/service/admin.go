// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "zeal_be/api/admin/v1"
)

type (
	IAdmin interface {
		CheckAdmin(ctx context.Context) int
		GetUserList(ctx context.Context, in v1.ULReqInfo) (stat int, userList []v1.UserInfo, count int)
		GetUserPrivs(ctx context.Context, t_email string) (privs []string)
		UpdateUserPrivs(ctx context.Context, t_email string, privs []string)
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
