package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/cnt"
	"focus/app/model"
	"focus/app/service"
)

// 问答管理
var Ask = askApi{}

type askApi struct{}

// @summary 展示问答首页
// @tags    前台-问答
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /ask [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Index(ctx context.Context, req *internal.ContentGetListReq) error {
	req.Type = cnt.ContentTypeAsk
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

// @summary 展示问答详情
// @tags    前台-问答
// @produce html
// @param   id path int false "问答ID"
// @router  /ask/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Detail(ctx context.Context, req *internal.ContentDetailReq) error {
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
		var (
			title = service.View.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			})
			breadCrumb = service.View.GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			})
		)
		service.View.Render(ctx, model.View{
			ContentType: cnt.ContentTypeAsk,
			Data:        getDetailRes,
			Title:       title,
			BreadCrumb:  breadCrumb,
		})
	}
	return nil
}
