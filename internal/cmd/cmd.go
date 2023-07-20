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

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"

	"focus-single/internal/service"

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

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					index.NewV1(),    // 首页
					login.NewV1(),    // 登录
					register.NewV1(), // 注册
					category.NewV1(), // 栏目
					topic.NewV1(),    // 主题
					ask.NewV1(),      // 问答
					article.NewV1(),  // 文章
					reply.NewV1(),    // 回复
					search.NewV1(),   // 搜索
					captcha.NewV1(),  // 验证码
					user.NewV1(),     // 用户
				)
				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						profile.NewV1(),  // 个人
						content.NewV1(),  // 内容
						interact.NewV1(), // 交互
						file.NewV1(),     // 文件
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
