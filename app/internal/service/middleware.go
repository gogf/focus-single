package service

import (
	"focus/app/internal/cnt"
	"focus/app/internal/model"
	"focus/library/response"
	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guid"
)

// 中间件管理服务
var Middleware = middlewareService{
	LoginUrl: "/login",
}

type middlewareService struct {
	LoginUrl string // 登录路由地址
}

// 返回处理中间件
func (s *middlewareService) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err  error
		res  interface{}
		code gcode.Code
	)
	res, err = r.GetHandlerResponse()
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), err.Error())
		} else {
			View.Render500(r.Context(), model.View{
				Error: err.Error(),
			})
		}
	} else {
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), "", res)
		} else {
			// 什么都不做，业务API自行处理模板渲染的成功逻辑。
		}
	}
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
		adminId := g.Cfg().MustGet(r.Context(), "setting.adminId", cnt.DefaultAdminId).Uint()
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
			response.JsonRedirectExit(r, 1, "", s.LoginUrl)
		} else {
			r.Response.RedirectTo(s.LoginUrl)
		}
	}
	r.Middleware.Next()
}

// 添加自定义RequestId到ctx中
func (s *middlewareService) RequestId(r *ghttp.Request) {
	requestId := r.GetCtxVar(cnt.ContextKeyRequestId).String()
	if requestId == "" {
		if v := r.Header.Get(cnt.ContextKeyRequestId); v != "" {
			requestId = v
		} else {
			requestId = guid.S()
		}
		r.SetCtxVar(cnt.ContextKeyRequestId, requestId)
		r.Response.Header().Set(cnt.ContextKeyRequestId, requestId)
	}
	r.Middleware.Next()
}
