package captcha

import (
	"context"

	"focus-single/api/captcha/v1"
	"focus-single/internal/consts"
	"focus-single/internal/service"
)

func (c *ControllerV1) CaptchaIndex(ctx context.Context, req *v1.CaptchaIndexReq) (res *v1.CaptchaIndexRes, err error) {
	err = service.Captcha().NewAndStore(ctx, consts.CaptchaDefaultName)
	return
}
