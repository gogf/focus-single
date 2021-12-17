package apiv1

import "github.com/gogf/gf/v2/frame/g"

type RegisterIndexReq struct {
	g.Meta `path:"/register" method:"get" summary:"展示注册页面" tags:"注册"`
}
type RegisterIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type RegisterDoReq struct {
	g.Meta   `path:"/register" method:"post" summary:"执行注册请求" tags:"注册"`
	Passport string `json:"passport" v:"required#请输入账号" dc:"账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
	Nickname string `json:"nickname" v:"required#请输入昵称" dc:"昵称"`
	Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}
type RegisterDoRes struct{}
