package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/cnt"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// 登录管理
var Login = loginApi{}

type loginApi struct{}

// @summary 展示登录页面
// @tags    前台-登录
// @produce html
// @router  /login [GET]
// @success 200 {string} html "页面HTML"
func (a *loginApi) Index(ctx context.Context) {
	service.View.Render(ctx, model.View{})
}

// @summary 提交登录
// @description 前面5次不需要验证码，同一个IP登录失败5次之后将会启用验证码校验。
// @description 注意提交的密码是明文。
// @description 登录成功后前端引导页面跳转。
// @tags    前台-登录
// @produce json
// @param   passport    formData string true "账号"
// @param   password    formData string true "密码"
// @param   verify_code formData string false "验证码"
// @router  /login/do [POST]
// @success 200 {object} response.JsonRes "执行结果"
func (a *loginApi) Do(ctx context.Context, req *internal.UserLoginReq) (string, error) {
	if !service.Captcha.VerifyAndClear(g.RequestFromCtx(ctx), cnt.CaptchaDefaultName, req.Captcha) {
		return "", gerror.NewCode(gerror.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	if err := service.User.Login(ctx, req.UserLoginInput); err != nil {
		return "", err
	} else {
		// 识别并跳转到登录前页面
		loginReferer := service.Session.GetLoginReferer(ctx)
		if loginReferer != "" {
			_ = service.Session.RemoveLoginReferer(ctx)
		}
		return loginReferer, nil
	}
}
