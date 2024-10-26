package auth

import (
	"github.com/devpenguin/go-ecommerce/internal/service"
	"github.com/devpenguin/go-ecommerce/internal/vo"
	"github.com/devpenguin/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

var UserAuth = new(sUserAuth)
type sUserAuth struct{}

func (c *sUserAuth) Register(ctx *gin.Context) {
    // Implement logic for register
	req := vo.RegisterUser_Request{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.CODE_BAD_GATEWAY, "Fail to extract data")
		return
	}
    code, res, err := service.UserAuth().Register(ctx, &req)
    if err != nil {
        response.ErrorResponse(ctx, response.CODE_INTERNAL_SERVER_ERROR, err.Error())
        return
    }

    response.SuccessResponse(ctx, int(code), "Register success", res)
}
