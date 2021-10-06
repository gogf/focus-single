package api

import (
	"github.com/gogf/gf/frame/g"
)

// 验证码展示
type CaptchaIndexReq struct {
	g.Meta `method:"get" summary:"获取默认的验证码" description:"注意直接返回的是图片二进制内容" tags:"验证码"`
}
type CaptchaIndexRes struct {
	g.Meta `mime:"png" description:"验证码二进制内容" `
}
