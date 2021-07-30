package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/cnt"
	"focus/app/model"
	"focus/app/service"
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
func (a *articleApi) Index(ctx context.Context, req *internal.ContentGetListReq) error {
	req.Type = cnt.ContentTypeArticle
	if getListRes, err := service.Content.GetList(ctx, req.ContentGetListInput); err != nil {
		return err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: req.Type,
				CategoryId:  req.CategoryId,
			}),
		})
	}
	return nil
}

// @summary 展示文章详情
// @tags    前台-文章
// @produce html
// @param   id path int false "文章ID"
// @router  /article/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Detail(ctx context.Context, req *internal.ContentDetailReq) error {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return err
	} else {
		if getDetailRes == nil {
			service.View.Render404(ctx)
			return nil
		}
		if err := service.Content.AddViewCount(ctx, req.Id, 1); err != nil {
			return err
		}
		service.View.Render(ctx, model.View{
			ContentType: cnt.ContentTypeArticle,
			Data:        getDetailRes,
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			}),
			BreadCrumb: service.View.GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			}),
		})
	}
	return nil
}
