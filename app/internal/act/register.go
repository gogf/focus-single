package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/cnt"
	"focus/app/internal/model"
	"focus/app/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// 注册控制器
	Register = registerAct{}
)

type registerAct struct{}

func (a *registerAct) Index(ctx context.Context, req *api.RegisterIndexReq) (res *api.RegisterIndexRes, err error) {
	service.View.Render(ctx, model.View{})
	return
}

func (a *registerAct) Do(ctx context.Context, req *api.UserRegisterReq) (res *api.UserRegisterRes, err error) {
	if !service.Captcha.VerifyAndClear(g.RequestFromCtx(ctx), cnt.CaptchaDefaultName, req.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	if err = service.User.Register(ctx, req.UserRegisterInput); err != nil {
		return
	}

	// 自动登录
	err = service.User.Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
