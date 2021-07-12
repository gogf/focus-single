package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 注册控制器
var Register = registerApi{}

type registerApi struct{}

// @summary 展示注册页面
// @tags    前台-注册
// @produce html
// @router  /register [GET]
// @success 200 {string} html "页面HTML"
func (a *registerApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{})
}

// @summary 执行注册提交处理
// @description 注意提交的密码是明文。
// @description 注册成功后自动登录。前端页面引导跳转
// @tags    前台-注册
// @produce json
// @param   entity body model.UserRegisterReq true "请求参数" required
// @router  /register/do [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *registerApi) Do(r *ghttp.Request) {
	var (
		req *model.UserRegisterReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if !service.Captcha.VerifyAndClear(r, model.CaptchaDefaultName, req.Captcha) {
		response.JsonExit(r, 1, "请输入正确的验证码")
	}
	if err := service.User.Register(r.Context(), req.UserRegisterInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		// 自动登录
		err := service.User.Login(r.Context(), model.UserLoginInput{
			Passport: req.Passport,
			Password: req.Password,
		})
		if err != nil {
			response.JsonExit(r, 1, err.Error())
		}
		response.JsonExit(r, 0, "")
	}
}
