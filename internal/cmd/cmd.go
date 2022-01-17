package cmd

import (
	"context"

	"focus-single/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/protocol/goai"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"

	"focus-single/internal/consts"
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
					service.View().Render401(r.Context())
				}
			})
			s.BindStatusHandler(403, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render403(r.Context())
				}
			})
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render404(r.Context())
				}
			})
			s.BindStatusHandler(500, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render500(r.Context())
				}
			})

			// 前台系统路由注册
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Index,    // 首页
					controller.Login,    // 登录
					controller.Register, // 注册
					controller.Category, // 栏目
					controller.Topic,    // 主题
					controller.Ask,      // 问答
					controller.Article,  // 文章
					controller.Reply,    // 回复
					controller.Search,   // 搜索
					controller.Captcha,  // 验证码
					controller.User,     // 用户
				)
				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.Profile,  // 个人
						controller.Content,  // 内容
						controller.Interact, // 交互
						controller.File,     // 文件
					)
				})
			})
			// 自定义丰富文档
			enhanceOpenAPIDoc(s)
			// 启动Http Server
			s.Run()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `Focus Project`
	openapi.Info.Description = ``

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		{Name: consts.OpenAPITagNameIndex},
		{Name: consts.OpenAPITagNameLogin},
		{Name: consts.OpenAPITagNameRegister},
		{Name: consts.OpenAPITagNameArticle},
		{Name: consts.OpenAPITagNameTopic},
		{Name: consts.OpenAPITagNameAsk},
		{Name: consts.OpenAPITagNameReply},
		{Name: consts.OpenAPITagNameContent},
		{Name: consts.OpenAPITagNameSearch},
		{Name: consts.OpenAPITagNameInteract},
		{Name: consts.OpenAPITagNameCategory},
		{Name: consts.OpenAPITagNameProfile},
		{Name: consts.OpenAPITagNameUser},
		{Name: consts.OpenAPITagNameMess},
	}
}
