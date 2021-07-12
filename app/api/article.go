package api

import (
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 文章管理
var Article = articleApi{}

type articleApi struct{}

// @summary 展示文章首页
// @tags    前台-文章
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /article [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Index(r *ghttp.Request) {
	var (
		req *internal.ContentGetListReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	req.Type = model.ContentTypeArticle
	if getListRes, err := service.Content.GetList(r.Context(), req.ContentGetListInput); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title: service.View.GetTitle(r.Context(), &model.ViewGetTitleInput{
				ContentType: req.Type,
				CategoryId:  req.CategoryId,
			}),
		})
	}
}

// @summary 展示文章详情
// @tags    前台-文章
// @produce html
// @param   id path int false "文章ID"
// @router  /article/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Detail(r *ghttp.Request) {
	var (
		req *internal.ContentDetailReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getDetailRes, err := service.Content.GetDetail(r.Context(), req.Id); err != nil {
		service.View.Render500(r)
	} else {
		if getDetailRes == nil {
			service.View.Render404(r)
		}
		if err := service.Content.AddViewCount(r.Context(), req.Id, 1); err != nil {
			service.View.Render500(r, model.View{
				Error: err.Error(),
			})
		}
		service.View.Render(r, model.View{
			ContentType: model.ContentTypeArticle,
			Data:        getDetailRes,
			Title: service.View.GetTitle(r.Context(), &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			}),
			BreadCrumb: service.View.GetBreadCrumb(r.Context(), &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			}),
		})
	}
}
