package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 验证码展示
type CaptchaIndexReq struct {
	g.Meta `method:"get" summary:"获取默认的验证码" dc:"注意直接返回的是图片二进制内容" tags:"验证码"`
}
type CaptchaIndexRes struct {
	g.Meta `mime:"png" dc:"验证码二进制内容" `
}
