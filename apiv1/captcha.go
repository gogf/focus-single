package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaIndexReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"工具" summary:"获取默认的验证码" dc:"注意直接返回的是图片二进制内容"`
}
type CaptchaIndexRes struct {
	g.Meta `mime:"png" dc:"验证码二进制内容" `
}
