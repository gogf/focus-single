package controller

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Index(ctx context.Context, req *apiv1.LoginIndexReq) (res *apiv1.LoginIndexRes, err error) {
	service.View().Render(ctx, model.View{})
	return
}

func (a *cLogin) Login(ctx context.Context, req *apiv1.LoginDoReq) (res *apiv1.LoginDoRes, err error) {
	res = &apiv1.LoginDoRes{}
	if !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), consts.CaptchaDefaultName, req.Captcha) {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	if err = service.User().Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	}); err != nil {
		return
	} else {
		// 识别并跳转到登录前页面
		loginReferer := service.Session().GetLoginReferer(ctx)
		if loginReferer != "" {
			_ = service.Session().RemoveLoginReferer(ctx)
		}
		res.Referer = loginReferer
		return
	}
}
