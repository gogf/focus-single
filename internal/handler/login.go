package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// 登录管理
	Login = handlerLogin{}
)

type handlerLogin struct{}

func (a *handlerLogin) Index(ctx context.Context, req *apiv1.LoginIndexReq) (res *apiv1.LoginIndexRes, err error) {
	service.View.Render(ctx, model.View{})
	return
}

func (a *handlerLogin) Do(ctx context.Context, req *apiv1.UserLoginReq) (res *apiv1.UserLoginRes, err error) {
	res = &apiv1.UserLoginRes{}
	//if !service.Captcha.VerifyAndClear(g.RequestFromCtx(ctx), consts.CaptchaDefaultName, req.Captcha) {
	//	return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	//}
	if err = service.User.Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	}); err != nil {
		return
	} else {
		// 识别并跳转到登录前页面
		loginReferer := service.Session.GetLoginReferer(ctx)
		if loginReferer != "" {
			_ = service.Session.RemoveLoginReferer(ctx)
		}
		res.Referer = loginReferer
		return
	}
}
