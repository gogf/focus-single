package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/service"
)

// 图形验证码
var Captcha = cCaptcha{}

type cCaptcha struct{}

func (a *cCaptcha) Index(ctx context.Context, req *v1.CaptchaIndexReq) (res *v1.CaptchaIndexRes, err error) {
	err = service.Captcha().NewAndStore(ctx, consts.CaptchaDefaultName)
	return
}
