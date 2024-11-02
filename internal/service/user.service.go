package service

import (
	"context"

	"github.com/devpenguin/go-ecommerce/internal/vo"
)

type (
	IUserAuth interface {
		Login(ctx context.Context, in *vo.LoginInput) (codeResult int, out vo.LoginOutput, err error)
		Register(ctx context.Context, in *vo.RegisterUser_Request) (int32, interface{}, error)
		VerifyOTP(ctx context.Context, in *vo.VerifyInput) (vo.VerifyOTPOutput, error)
		UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error)
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