// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISystem interface {
		// 发送验证码
		EmailCodeSend(ctx context.Context, to string) (send bool, ttl int64)
		// 验证
		EmailCodeVerify(ctx context.Context, to string, code string) bool
		EmailCodeSend2(ctx context.Context, to string) (send bool, ttl int64)
		// 获取图形验证码
		GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error)
		// 图形验证码验证
		VerifyCaptcha(ctx context.Context, idKey string, captcha string) bool
		// 登录功能
		Login(ctx context.Context, email string, password string) (stat int, tokenString string)
		UpdateIPAddress(ctx context.Context, email string)
		// 注册
		Register(ctx context.Context, username string, email string, password string) int
		UpdatePwd(ctx context.Context, email string, password string) int
	}
)

var (
	localSystem ISystem
)

func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

func RegisterSystem(i ISystem) {
	localSystem = i
}
