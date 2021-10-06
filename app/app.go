package app

import (
	_ "focus/app/internal/packed"

	"focus/app/internal/act"
	"focus/app/internal/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmode"
)

// 应用启动
func Run() {
	var (
		ctx = gctx.New()
	)
	// 绑定Swagger Plugin
	s := g.Server()

	// 静态目录设置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		g.Log().Fatal(ctx, "文件上传配置路径不能为空")
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
			"/":             act.Index,          // 首页
			"/login":        act.Login,          // 登录
			"/register":     act.Register,       // 注册
			"/category":     act.Category,       // 栏目
			"/topic":        act.Topic,          // 主题
			"/topic/{id}":   act.Topic.Detail,   // 主题 - 详情
			"/ask":          act.Ask,            // 问答
			"/ask/{id}":     act.Ask.Detail,     // 问答 - 详情
			"/article":      act.Article,        // 文章
			"/article/{id}": act.Article.Detail, // 文章 - 详情
			"/reply":        act.Reply,          // 回复
			"/search":       act.Search,         // 搜索
			"/captcha":      act.Captcha,        // 验证码
			"/user/{id}":    act.User.Index,     // 用户 - 主页
		})
		// 权限控制路由
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALLMap(g.Map{
				"/user":     act.User,     // 用户
				"/content":  act.Content,  // 内容
				"/interact": act.Interact, // 交互
				"/file":     act.File,     // 文件
			})
		})
	})

	// 启动Http Server
	s.Run()
}
