package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

// Topic 主题管理
var Topic = cTopic{}

type cTopic struct{}

func (a *cTopic) Index(ctx context.Context, req *v1.TopicIndexReq) (res *v1.TopicIndexRes, err error) {
	req.Type = consts.ContentTypeTopic
	out, err := service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	title := service.View().GetTitle(ctx, &model.ViewGetTitleInput{
		ContentType: req.Type,
		CategoryId:  req.CategoryId,
	})
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
		Data:        out,
		Title:       title,
	})
	return
}

func (a *cTopic) Detail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error) {
	out, err := service.Content().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		service.View().Render404(ctx)
		return
	}
	err = service.Content().AddViewCount(ctx, req.Id, 1)
	service.View().Render(ctx, model.View{
		ContentType: consts.ContentTypeTopic,
		Data:        out,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
			CurrentName: out.Content.Title,
		}),
		BreadCrumb: service.View().GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
			ContentId:   out.Content.Id,
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
		}),
	})
	return
}
