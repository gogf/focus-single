package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"

	"focus-single/internal/handler"
	"focus-single/internal/service"
	"focus-single/utility/response"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start focus server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			// OpenApi自定义信息
			oai.Info.Title = `API Reference`
			oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

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
					service.Middleware.Ctx,
					service.Middleware.ResponseHandler,
				)
				group.Bind(
					handler.Index,    // 首页
					handler.Login,    // 登录
					handler.Register, // 注册
					handler.Category, // 栏目
					handler.Topic,    // 主题
					handler.Ask,      // 问答
					handler.Article,  // 文章
					handler.Reply,    // 回复
					handler.Search,   // 搜索
					handler.Captcha,  // 验证码
					handler.User,     // 用户
				)
				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware.Auth)
					group.Bind(
						handler.Profile,  // 个人
						handler.Content,  // 内容
						handler.Interact, // 交互
						handler.File,     // 文件
					)
				})
			})

			// 启动Http Server
			s.Run()
			return
		},
	}
)
