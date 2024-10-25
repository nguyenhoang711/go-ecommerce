package service

import (
	"context"

	"github.com/devpenguin/go-ecommerce/internal/vo"
)

type (
	ISystemAPI interface {
		CMS_PING(ctx context.Context, in *vo.CMSPing_Request) (codeResult int, data interface{}, err error)
	}
)

var (
	localSystemAPI ISystemAPI
)

func SystemAPI() ISystemAPI {
	if localSystemAPI == nil {
		panic("Implement localSystemAPI not found for interface ISystemAPI")
	}
	return localSystemAPI
}

func InitSystemAPI(i ISystemAPI) {
	localSystemAPI = i
}