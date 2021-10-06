package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/cnt"
	"focus/app/internal/service"
)

var (
	// 图形验证码
	Captcha = captchaAct{}
)

type captchaAct struct{}

func (a *captchaAct) Index(ctx context.Context, req *api.CaptchaIndexReq) (res *api.CaptchaIndexRes, err error) {
	err = service.Captcha.NewAndStore(ctx, cnt.CaptchaDefaultName)
	return
}
