package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/cnt"
	"focus/app/internal/model"
	"focus/app/internal/service"
	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var (
	// 注册控制器
	Register = registerAct{}
)

type registerAct struct{}

// @summary 展示注册页面
// @tags    前台-注册
// @produce html
// @router  /register [GET]
// @success 200 {string} html "页面HTML"
func (a *registerAct) Index(ctx context.Context, req *api.RegisterIndexReq) (res *api.RegisterIndexRes, err error) {
	service.View.Render(ctx, model.View{})
	return
}

// @summary 执行注册提交处理
// @description 注意提交的密码是明文。
// @description 注册成功后自动登录。前端页面引导跳转
// @tags    前台-注册
// @produce json
// @param   entity body internal.UserRegisterReq true "请求参数" required
// @router  /register/do [POST]
// @success 200 {object} response.JsonRes "请求结果"
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
