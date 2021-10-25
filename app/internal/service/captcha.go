package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mojocn/base64Captcha"
)

// 验证码管理服务
var (
	Captcha = captchaService{}
)

type captchaService struct{}

var (
	captchaStore  = base64Captcha.DefaultMemStore
	captchaDriver = newDriver()
)

func newDriver() *base64Captcha.DriverString {
	driver := &base64Captcha.DriverString{
		Height:          44,
		Width:           126,
		NoiseCount:      5,
		ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "1234567890",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

// 创建验证码，直接输出验证码图片内容到HTTP Response.
func (s *captchaService) NewAndStore(ctx context.Context, name string) error {
	var (
		request = g.RequestFromCtx(ctx)
		captcha = base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	)
	_, content, answer := captcha.Driver.GenerateIdQuestionAnswer()
	item, _ := captcha.Driver.DrawCaptcha(content)
	captchaStoreKey := guid.S()
	request.Session.Set(name, captchaStoreKey)
	captcha.Store.Set(captchaStoreKey, answer)
	_, err := item.WriteTo(request.Response.Writer)
	return err
}

// 校验验证码，并清空缓存的验证码信息
func (s *captchaService) VerifyAndClear(r *ghttp.Request, name string, value string) bool {
	defer r.Session.Remove(name)
	captchaStoreKey := r.Session.MustGet(name).String()
	return captchaStore.Verify(captchaStoreKey, value, true)
}
