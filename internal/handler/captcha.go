package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/service"
)

var (
	// 图形验证码
	Captcha = handlerCaptcha{}
)

type handlerCaptcha struct{}

func (a *handlerCaptcha) Index(ctx context.Context, req *apiv1.CaptchaIndexReq) (res *apiv1.CaptchaIndexRes, err error) {
	err = service.Captcha.NewAndStore(ctx, consts.CaptchaDefaultName)
	return
}
