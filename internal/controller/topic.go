package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

// Topic 主题管理
var Topic = cTopic{}

type cTopic struct{}

func (a *cTopic) Index(ctx context.Context, req *v1.TopicIndexReq) (res *v1.TopicIndexRes, err error) {
	req.Type = consts.ContentTypeTopic
	out, err := content.GetList(ctx, content.GetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	title := view.GetTitle(ctx, &view.GetTitleInput{
		ContentType: req.Type,
		CategoryId:  req.CategoryId,
	})
	view.Render(ctx, view.View{
		ContentType: req.Type,
		Data:        out,
		Title:       title,
	})
	return
}

func (a *cTopic) Detail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error) {
	out, err := content.GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		view.Render404(ctx)
		return
	}
	err = content.AddViewCount(ctx, req.Id, 1)
	view.Render(ctx, view.View{
		ContentType: consts.ContentTypeTopic,
		Data:        out,
		Title: view.GetTitle(ctx, &view.GetTitleInput{
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
			CurrentName: out.Content.Title,
		}),
		BreadCrumb: view.GetBreadCrumb(ctx, &view.GetBreadCrumbInput{
			ContentId:   out.Content.Id,
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
		}),
	})
	return
}
