package syst

import (
	"github.com/devpenguin/go-ecommerce/internal/service"
	"github.com/devpenguin/go-ecommerce/internal/vo"
	"github.com/devpenguin/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

var SystemAPI = new(sSystemAPI)
type sSystemAPI struct{}

// @Summary cms Server info
// @Description Lấy về thông tin cơ bản của server , như giờ hệ thống , tên server , version hiện tại của hệ thống
// @Accept  json
// @Produce  json
// @Tags INFO
// @Param 	request	body vo.CMSPing_Request true	"info"
// @Success		200	{object} 	vo.CMSPing_Reply
// @Router /syst/system-info [post]
func (s *sSystemAPI) Ping(ctx *gin.Context) {
	req := vo.CMSPing_Request{}
	code, data, err := service.SystemAPI().CMS_PING(ctx, &req)
	if err != nil {
		response.ErrorResponse(ctx, code, "System info fail")
		return
	}
	response.SuccessResponse(ctx, code, "System info show", data)
}
