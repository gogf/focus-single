package app

import (
	_ "focus/packed"

	"focus/app/api"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmode"
	"github.com/gogf/swagger"
)

// 应用启动
func Run() {
	// 绑定Swagger Plugin
	s := g.Server()
	s.Plugin(&swagger.Swagger{})

	// 静态目录设置
	uploadPath := g.Cfg().GetString("upload.path")
	if uploadPath == "" {
		g.Log().Fatal("文件上传配置路径不能为空")
	}
	s.AddStaticPath("/upload", uploadPath)

	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	// 前台系统自定义错误页面
	s.BindStatusHandler(401, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			service.View.Render401(r.Context())
		}
	})
	s.BindStatusHandler(403, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			service.View.Render403(r.Context())
		}
	})
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			service.View.Render404(r.Context())
		}
	})
	s.BindStatusHandler(500, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			service.View.Render500(r.Context())
		}
	})

	// 前台系统路由注册
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.RequestId,
			service.Middleware.Ctx,
			service.Middleware.ResponseHandler,
		)
		group.ALLMap(g.Map{
			"/":            api.Index,          // 首页
			"/login":       api.Login,          // 登录
			"/register":    api.Register,       // 注册
			"/category":    api.Category,       // 栏目
			"/topic":       api.Topic,          // 主题
			"/topic/:id":   api.Topic.Detail,   // 主题 - 详情
			"/ask":         api.Ask,            // 问答
			"/ask/:id":     api.Ask.Detail,     // 问答 - 详情
			"/article":     api.Article,        // 文章
			"/article/:id": api.Article.Detail, // 文章 - 详情
			"/reply":       api.Reply,          // 回复
			"/search":      api.Search,         // 搜索
			"/captcha":     api.Captcha,        // 验证码
			"/user/:id":    api.User.Index,     // 用户 - 主页
		})
		// 权限控制路由
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALLMap(g.Map{
				"/user":     api.User,     // 用户
				"/content":  api.Content,  // 内容
				"/interact": api.Interact, // 交互
				"/file":     api.File,     // 文件
			})
		})
	})

	// 启动Http Server
	s.Run()
}
