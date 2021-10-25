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
	// 登录管理
	Login = loginAct{}
)

type loginAct struct{}

func (a *loginAct) Index(ctx context.Context, req *api.LoginIndexReq) (res *api.LoginIndexRes, err error) {
	service.View.Render(ctx, model.View{})
	return
}

func (a *loginAct) Do(ctx context.Context, req *api.UserLoginReq) (res *api.UserLoginRes, err error) {
	res = &api.UserLoginRes{}
	if !service.Captcha.VerifyAndClear(g.RequestFromCtx(ctx), cnt.CaptchaDefaultName, req.Captcha) {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	if err = service.User.Login(ctx, req.UserLoginInput); err != nil {
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
