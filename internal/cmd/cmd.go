package cmd

import (
	"context"

	"focus-single/internal/controller/article"
	"focus-single/internal/controller/ask"
	"focus-single/internal/controller/captcha"
	"focus-single/internal/controller/category"
	"focus-single/internal/controller/content"
	"focus-single/internal/controller/file"
	"focus-single/internal/controller/index"
	"focus-single/internal/controller/interact"
	"focus-single/internal/controller/login"
	"focus-single/internal/controller/profile"
	"focus-single/internal/controller/register"
	"focus-single/internal/controller/reply"
	"focus-single/internal/controller/search"
	"focus-single/internal/controller/topic"
	"focus-single/internal/controller/user"
	"focus-single/internal/service/middleware"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/protocol/goai"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"

	"focus-single/internal/consts"
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
					view.Render401(r.Context())
				}
			})
			s.BindStatusHandler(403, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					view.Render403(r.Context())
				}
			})
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					view.Render404(r.Context())
				}
			})
			s.BindStatusHandler(500, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					view.Render500(r.Context())
				}
			})

			// 前台系统路由注册
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middleware.Ctx,
					middleware.ResponseHandler,
				)
				group.Bind(
					index.New(),    // 首页
					login.New(),    // 登录
					register.New(), // 注册
					category.New(), // 栏目
					topic.New(),    // 主题
					ask.New(),      // 问答
					article.New(),  // 文章
					reply.New(),    // 回复
					search.New(),   // 搜索
					captcha.New(),  // 验证码
					user.New(),     // 用户
				)
				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.Bind(
						profile.New(),  // 个人
						content.New(),  // 内容
						interact.New(), // 交互
						file.New(),     // 文件
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
