package middleware

import (
	"focus-single/internal/service/bizctx"
	"focus-single/internal/service/session"
	"focus-single/internal/service/user"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/utility/response"
)

const (
	LoginUrl = "/login"
)

// 返回处理中间件
func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), err.Error())
		} else {
			view.Render500(r.Context(), model.View{
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
func Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	bizctx.Init(r, customCtx)
	if userEntity := session.GetUser(r.Context()); userEntity.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:       userEntity.Id,
			Passport: userEntity.Passport,
			Nickname: userEntity.Nickname,
			Avatar:   userEntity.Avatar,
			IsAdmin:  user.IsAdmin(r.Context(), userEntity.Id),
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
func Auth(r *ghttp.Request) {
	sessionUser := session.GetUser(r.Context())
	if sessionUser.Id == 0 {
		_ = session.SetNotice(r.Context(), &model.SessionNotice{
			Type:    consts.SessionNoticeTypeWarn,
			Content: "未登录或会话已过期，请您登录后再继续",
		})
		// 只有GET请求才支持保存当前URL，以便后续登录后再跳转回来。
		if r.Method == "GET" {
			_ = session.SetLoginReferer(r.Context(), r.GetUrl())
		}
		// 根据当前请求方式执行不同的返回数据结构
		if r.IsAjaxRequest() {
			response.JsonRedirectExit(r, 1, "", LoginUrl)
		} else {
			r.Response.RedirectTo(LoginUrl)
		}
	}
	r.Middleware.Next()
}
