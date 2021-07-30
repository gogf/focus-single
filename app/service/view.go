package service

import (
	"context"
	"fmt"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gmode"
)

// 视图管理服务
var View = viewService{}

type viewService struct{}

// 前台系统-获取面包屑列表
func (s *viewService) GetBreadCrumb(ctx context.Context, input *model.ViewGetBreadCrumbInput) []model.ViewBreadCrumb {
	breadcrumb := []model.ViewBreadCrumb{
		{Name: "首页", Url: "/"},
	}
	var uriPrefix string
	if input.ContentType != "" {
		uriPrefix = "/" + input.ContentType
		topMenuItem, _ := Menu.GetTopMenuByUrl(ctx, uriPrefix)
		if topMenuItem != nil {
			breadcrumb = append(breadcrumb, model.ViewBreadCrumb{
				Name: topMenuItem.Name,
				Url:  topMenuItem.Url,
			})
		}
	}
	if uriPrefix != "" && input.CategoryId > 0 {
		category, _ := Category.GetItem(ctx, input.CategoryId)
		if category != nil {
			breadcrumb = append(breadcrumb, model.ViewBreadCrumb{
				Name: category.Name,
				Url:  fmt.Sprintf(`%s?cate=%d`, uriPrefix, category.Id),
			})
		}
	}
	if input.ContentId > 0 {
		breadcrumb = append(breadcrumb, model.ViewBreadCrumb{
			Name: "内容详情",
		})
	}
	return breadcrumb
}

// 前台系统-获取标题
func (s *viewService) GetTitle(ctx context.Context, input *model.ViewGetTitleInput) string {
	var (
		titleArray []string
		uriPrefix  string
	)
	if input.CurrentName != "" {
		titleArray = append(titleArray, input.CurrentName)
	}
	if input.CategoryId > 0 {
		category, _ := Category.GetItem(ctx, input.CategoryId)
		if category != nil {
			titleArray = append(titleArray, category.Name)
		}
	}
	if input.ContentType != "" {
		uriPrefix = "/" + input.ContentType
		topMenuItem, _ := Menu.GetTopMenuByUrl(ctx, uriPrefix)
		if topMenuItem != nil {
			titleArray = append(titleArray, topMenuItem.Name)
		}
	}
	return gstr.Join(titleArray, " - ")
}

// 渲染指定模板页面
func (s *viewService) RenderTpl(ctx context.Context, tpl string, data ...model.View) {
	var (
		viewObj  = model.View{}
		viewData = make(g.Map)
		request  = g.RequestFromCtx(ctx)
	)
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.Title == "" {
		viewObj.Title = g.Cfg().GetString(`setting.title`)
	} else {
		viewObj.Title = viewObj.Title + ` - ` + g.Cfg().GetString(`setting.title`)
	}
	if viewObj.Keywords == "" {
		viewObj.Keywords = g.Cfg().GetString(`setting.keywords`)
	}
	if viewObj.Description == "" {
		viewObj.Description = g.Cfg().GetString(`setting.description`)
	}
	// 去掉空数据
	viewData = gconv.Map(viewObj)
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}
	// 内置对象
	viewData["BuildIn"] = &viewBuildIn{httpRequest: request}
	// 内容模板
	if viewData["MainTpl"] == nil {
		viewData["MainTpl"] = s.getDefaultMainTpl(ctx)
	}
	// 提示信息
	if notice, _ := Session.GetNotice(ctx); notice != nil {
		_ = Session.RemoveNotice(ctx)
		viewData["Notice"] = notice
	}
	// 渲染模板
	_ = request.Response.WriteTpl(tpl, viewData)
	// 开发模式下，在页面最下面打印所有的模板变量
	if gmode.IsDevelop() {
		_ = request.Response.WriteTplContent(`{{dump .}}`, viewData)
	}
}

// 渲染默认模板页面
func (s *viewService) Render(ctx context.Context, data ...model.View) {
	s.RenderTpl(ctx, g.Cfg().GetString("viewer.indexLayout"), data...)
}

// 跳转中间页面
func (s *viewService) Render302(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "页面跳转中"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/302.html"
	s.Render(ctx, view)
}

// 401页面
func (s *viewService) Render401(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/401.html"
	s.Render(ctx, view)
}

// 403页面
func (s *viewService) Render403(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/403.html"
	s.Render(ctx, view)
}

// 404页面
func (s *viewService) Render404(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "资源不存在"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/404.html"
	s.Render(ctx, view)
}

// 500页面
func (s *viewService) Render500(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "请求执行错误"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/500.html"
	s.Render(ctx, view)
}

// 获取视图存储目录
func (s *viewService) getViewFolderName() string {
	return gstr.Split(g.Cfg().GetString("viewer.indexLayout"), "/")[0]
}

// 获取自动设置的MainTpl
func (s *viewService) getDefaultMainTpl(ctx context.Context) string {
	var (
		viewFolderPrefix = s.getViewFolderName()
		urlPathArray     = gstr.SplitAndTrim(g.RequestFromCtx(ctx).URL.Path, "/")
		mainTpl          string
	)
	if len(urlPathArray) > 0 && urlPathArray[0] == viewFolderPrefix {
		urlPathArray = urlPathArray[1:]
	}
	switch {
	case len(urlPathArray) == 2:
		// 如果2级路由为数字，那么为模块的详情页面，那么路由固定为/xxx/detail。
		// 如果需要定制化内容模板，请在具体路由方法中设置MainTpl。
		if gstr.IsNumeric(urlPathArray[1]) {
			urlPathArray[1] = "detail"
		}
		mainTpl = viewFolderPrefix + "/" + gfile.Join(urlPathArray[0], urlPathArray[1]) + ".html"
	case len(urlPathArray) == 1:
		mainTpl = viewFolderPrefix + "/" + urlPathArray[0] + "/index.html"
	default:
		// 默认首页内容
		mainTpl = viewFolderPrefix + "/index/index.html"
	}
	return gstr.TrimLeft(mainTpl, "/")
}
