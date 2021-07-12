package service

import (
	"focus/app/cnt"
	"focus/app/model"
	"focus/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 中间件管理服务
var Middleware = middlewareService{
	loginUrl: "/login",
}

type middlewareService struct {
	loginUrl string // 登录路由地址
}

// 获取用户登录URL
func (s *middlewareService) GetLoginUrl() string {
	return s.loginUrl
}

// 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	Context.Init(r, customCtx)
	if userEntity := Session.GetUser(r.Context()); userEntity != nil {

		adminId := g.Cfg().GetUint("setting.adminId", cnt.DefaultAdminId)
		customCtx.User = &model.ContextUser{
			Id:       userEntity.Id,
			Passport: userEntity.Passport,
			Nickname: userEntity.Nickname,
			Avatar:   userEntity.Avatar,
			IsAdmin:  userEntity.Id == adminId,
		}
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 前台系统权限控制，用户必须登录才能访问
func (s *middlewareService) Auth(r *ghttp.Request) {
	user := Session.GetUser(r.Context())
	if user == nil {
		Session.SetNotice(r.Context(), &model.SessionNotice{
			Type:    cnt.SessionNoticeTypeWarn,
			Content: "未登录或会话已过期，请您登录后再继续",
		})
		// 只有GET请求才支持保存当前URL，以便后续登录后再跳转回来。
		if r.Method == "GET" {
			Session.SetLoginReferer(r.Context(), r.GetUrl())
		}
		// 根据当前请求方式执行不同的返回数据结构
		if r.IsAjaxRequest() {
			response.JsonRedirectExit(r, 1, "", s.loginUrl)
		} else {
			r.Response.RedirectTo(s.loginUrl)
		}
	}
	r.Middleware.Next()
}
