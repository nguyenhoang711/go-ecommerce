package auth

import (
	"github.com/devpenguin/go-ecommerce/internal/service"
	"github.com/devpenguin/go-ecommerce/internal/vo"
	"github.com/devpenguin/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

var UserAuth = new(cUserAuth)
type cUserAuth struct{}

func (c *cUserAuth) Register(ctx *gin.Context) {
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

// @Summary auth user verify otp
// @Description When user is registered send otp to email
// @Accept  json
// @Produce  json
// @Tags auth user management
// @Param 	payload	body vo.VerifyInput true	"payload"
// @Success		200	{object} 	response.ResponseData
// @Failure		500	{object} 	response.ErrorResponseData
// @Router /auth/verify_account [post]
func (c *cUserAuth) VerifyOTP(ctx *gin.Context) {
	var params vo.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.CODE_BAD_REQUEST, response.BAD_REQUEST)
		return
	}
	result, err := service.UserAuth().VerifyOTP(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrIOTPInvalid, response.ErrMessageDict()(response.ErrIOTPInvalid))
	}
	response.SuccessResponse(ctx, response.CODE_OK, "Xác thực OTP thành công", result)
}
