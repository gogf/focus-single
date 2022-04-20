package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

// Ask 问答管理
var Ask = cAak{}

type cAak struct{}

func (a *cAak) Index(ctx context.Context, req *v1.AskIndexReq) (res *v1.AskIndexRes, err error) {
	req.Type = consts.ContentTypeAsk
	getListRes, err := content.GetList(ctx, content.GetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	view.Render(ctx, view.View{
		ContentType: req.Type,
		Data:        getListRes,
		Title: view.GetTitle(ctx, &view.GetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		}),
	})
	return
}

func (a *cAak) Detail(ctx context.Context, req *v1.AskDetailReq) (res *v1.AskDetailRes, err error) {
	getDetailRes, err := content.GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if getDetailRes == nil {
		view.Render404(ctx)
		return nil, nil
	}
	if err = content.AddViewCount(ctx, req.Id, 1); err != nil {
		return nil, err
	}
	var (
		title = view.GetTitle(ctx, &view.GetTitleInput{
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
			CurrentName: getDetailRes.Content.Title,
		})
		breadCrumb = view.GetBreadCrumb(ctx, &view.GetBreadCrumbInput{
			ContentId:   getDetailRes.Content.Id,
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
		})
	)
	view.Render(ctx, view.View{
		ContentType: consts.ContentTypeAsk,
		Data:        getDetailRes,
		Title:       title,
		BreadCrumb:  breadCrumb,
	})
	return
}
