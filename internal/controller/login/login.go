package login

import (
	"context"

	v1 "focus-single/api/v1/login"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service/captcha"
	"focus-single/internal/service/session"
	"focus-single/internal/service/user"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	view.Render(ctx, model.View{})
	return
}

func (c *controller) Login(ctx context.Context, req *v1.DoReq) (res *v1.DoRes, err error) {
	res = &v1.DoRes{}
	if !captcha.VerifyAndClear(g.RequestFromCtx(ctx), consts.CaptchaDefaultName, req.Captcha) {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	err = user.Login(ctx, model.UserLoginInput{
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
