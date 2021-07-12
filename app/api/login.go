package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 登录管理
var Login = loginApi{}

type loginApi struct{}

// @summary 展示登录页面
// @tags    前台-登录
// @produce html
// @router  /login [GET]
// @success 200 {string} html "页面HTML"
func (a *loginApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{})
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
func (a *loginApi) Do(r *ghttp.Request) {
	var (
		req *model.UserLoginReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if !service.Captcha.VerifyAndClear(r, model.CaptchaDefaultName, req.Captcha) {
		response.JsonExit(r, 1, "请输入正确的验证码")
	}
	if err := service.User.Login(r.Context(), req.UserLoginInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		// 识别并跳转到登录前页面
		loginReferer := service.Session.GetLoginReferer(r.Context())
		if loginReferer != "" {
			_ = service.Session.RemoveLoginReferer(r.Context())
		}
		response.JsonRedirectExit(r, 0, "", loginReferer)
	}
}
