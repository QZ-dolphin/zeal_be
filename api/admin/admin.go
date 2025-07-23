// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"zeal_be/api/admin/v1"
)

type IAdminV1 interface {
	GetSelfPrivs(ctx context.Context, req *v1.GetSelfPrivsReq) (res *v1.GetSelfPrivsRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	GetUserPrivs(ctx context.Context, req *v1.GetUserPrivsReq) (res *v1.GetUserPrivsRes, err error)
	UpdateUserPrivs(ctx context.Context, req *v1.UpdateUserPrivsReq) (res *v1.UpdateUserPrivsRes, err error)
}
