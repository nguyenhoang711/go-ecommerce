package impl

import (
	"context"
	"strconv"

	"github.com/devpenguin/go-ecommerce/global"
	"github.com/devpenguin/go-ecommerce/internal/vo"
	"github.com/devpenguin/go-ecommerce/pkg/response"
)

type sSystemAPI struct {}

func NewSystemAPIImpl() *sSystemAPI {
	return &sSystemAPI{}
}

// implement ISystemAPI interface here
func (s *sSystemAPI) CMS_PING(ctx context.Context, in *vo.CMSPing_Request) (int, *vo.CMSPing_Reply, error) {
	global.Logger.Info(in.Hello)
	reply := vo.CMSPing_Reply{}
	reply.Port = strconv.Itoa(global.Config.Server.Port)
	reply.Mode = global.Config.Server.Mode
	return response.CODE_OK, &reply, nil
}
