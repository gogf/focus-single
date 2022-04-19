package register

import (
	"context"

	v1 "focus-single/api/v1/register"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service/captcha"
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

func (c *controller) Register(ctx context.Context, req *v1.DoReq) (res *v1.DoRes, err error) {
	if !captcha.VerifyAndClear(g.RequestFromCtx(ctx), consts.CaptchaDefaultName, req.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	if err = user.Register(ctx, model.UserRegisterInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}); err != nil {
		return
	}

	// 自动登录
	err = user.Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
