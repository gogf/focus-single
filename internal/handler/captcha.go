package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/service"
)

// 图形验证码
var Captcha = hCaptcha{}

type hCaptcha struct{}

func (a *hCaptcha) Index(ctx context.Context, req *apiv1.CaptchaIndexReq) (res *apiv1.CaptchaIndexRes, err error) {
	err = service.Captcha().NewAndStore(ctx, consts.CaptchaDefaultName)
	return
}
