package api

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 图形验证码
var Captcha = captchaApi{}

type captchaApi struct{}

// @summary 获取默认的验证码
// @description 注意直接返回的是图片二进制内容。
// @tags    前台-验证码
// @produce png
// @router  /captcha [GET]
// @success 200 {file} body "验证码二进制内容"
func (a *captchaApi) Index(r *ghttp.Request) {
	service.Captcha.NewAndStore(r, model.CaptchaDefaultName)
}
