// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "zeal_be/api/social/v1"
)

type (
	ISocial interface {
		SearchId(ctx context.Context, in v1.SearchReqInfo) (user_list []v1.User_info, total int)
		IsFollowing(ctx context.Context, user_id int) bool
		IsFollower(ctx context.Context, user_id int) bool
		IsInSocial(ctx context.Context) bool
		CreateSocialAccount(ctx context.Context)
		IsInSocialv2(ctx context.Context, user_id int) bool
		CreateSocialAccountV2(ctx context.Context, user_id int)
		GetSocialInfo(ctx context.Context) (following int, followers int)
		ToFollowing(ctx context.Context, user_id int, user_name string, nick_name string, avatar_url string) bool
		NotFollowing(ctx context.Context, user_id int) bool
		GetFollowerList(ctx context.Context, pageSize int, pageIndex int) []v1.Userdata
		GetFollowingList(ctx context.Context, pageSize int, pageIndex int) []v1.Userdata
	}
)

var (
	localSocial ISocial
)

func Social() ISocial {
	if localSocial == nil {
		panic("implement not found for interface ISocial, forgot register?")
	}
	return localSocial
}

func RegisterSocial(i ISocial) {
	localSocial = i
}
