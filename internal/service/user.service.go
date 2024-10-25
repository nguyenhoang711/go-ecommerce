package service

import (
	"context"

	"github.com/devpenguin/go-ecommerce/internal/vo"
)

type (
	IUserAuth interface {
		Login(ctx context.Context) error
		Register(ctx context.Context, in *vo.RegisterUser_Request) (int, vo.RegisterUser_Reply, error)
		VerifyOTP(ctx context.Context) error
		UpdatePasswordRegister(ctx context.Context) error
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
	}
)

var (
	localUserAuth IUserAuth
	localUserInfo IUserInfo
)

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("Implement localUserInfo not found for interface IUserInfo")
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func UserAuth() IUserAuth {
	if localUserAuth == nil {
		panic("Implement localUserAuth not found for interface IUserAuth")
	}
	return localUserAuth
}

func InitUserAuth(i IUserAuth) {
	localUserAuth = i
}