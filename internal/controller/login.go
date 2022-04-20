package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/service/captcha"
	"focus-single/internal/service/session"
	"focus-single/internal/service/user"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Index(ctx context.Context, req *v1.LoginIndexReq) (res *v1.LoginIndexRes, err error) {
	view.Render(ctx, view.View{})
	return
}

func (a *cLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	if !captcha.VerifyAndClear(g.RequestFromCtx(ctx), consts.CaptchaDefaultName, req.Captcha) {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	err = user.Login(ctx, user.LoginInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	// 识别并跳转到登录前页面
	loginReferer := session.GetLoginReferer(ctx)
	if loginReferer != "" {
		_ = session.RemoveLoginReferer(ctx)
	}
	res.Referer = loginReferer
	return
}
