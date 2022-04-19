package captcha

import (
	"context"

	v1 "focus-single/api/v1/captcha"
	"focus-single/internal/consts"
	"focus-single/internal/service/captcha"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	err = captcha.NewAndStore(ctx, consts.CaptchaDefaultName)
	return
}
