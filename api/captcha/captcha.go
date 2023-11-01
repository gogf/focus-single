// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package captcha

import (
	"context"

	"focus-single/api/captcha/v1"
)

type ICaptchaV1 interface {
	CaptchaIndex(ctx context.Context, req *v1.CaptchaIndexReq) (res *v1.CaptchaIndexRes, err error)
}
