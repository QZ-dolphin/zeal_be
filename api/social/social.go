// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package social

import (
	"context"

	"zeal_be/api/social/v1"
)

type ISocialV1 interface {
	Follow(ctx context.Context, req *v1.FollowReq) (res *v1.FollowRes, err error)
	GetFollowerList(ctx context.Context, req *v1.GetFollowerListReq) (res *v1.GetFollowerListRes, err error)
	GetFollowingList(ctx context.Context, req *v1.GetFollowingListReq) (res *v1.GetFollowingListRes, err error)
	GetSocialInfo(ctx context.Context, req *v1.GetSocialInfoReq) (res *v1.GetSocialInfoRes, err error)
	NotFollow(ctx context.Context, req *v1.NotFollowReq) (res *v1.NotFollowRes, err error)
	SearchId(ctx context.Context, req *v1.SearchIdReq) (res *v1.SearchIdRes, err error)
}
