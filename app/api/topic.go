package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/cnt"
	"focus/app/model"
	"focus/app/service"
)

// 主题管理
var Topic = topicApi{}

type topicApi struct{}

// @summary 展示主题首页
// @tags    前台-主题
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /topic [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Index(ctx context.Context, req *internal.ContentGetListReq) error {
	req.Type = cnt.ContentTypeTopic
	if getListRes, err := service.Content.GetList(ctx, req.ContentGetListInput); err != nil {
		return err
	} else {
		title := service.View.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		})
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       title,
		})
		return nil
	}
}

// @summary 展示主题详情
// @tags    前台-主题
// @produce html
// @param   id path int false "主题ID"
// @router  /topic/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Detail(ctx context.Context, req *internal.ContentDetailReq) error {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return err
	} else {
		if getDetailRes == nil {
			service.View.Render404(ctx)
			return nil
		}
		service.Content.AddViewCount(ctx, req.Id, 1)
		service.View.Render(ctx, model.View{
			ContentType: cnt.ContentTypeTopic,
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
		return nil
	}
}
