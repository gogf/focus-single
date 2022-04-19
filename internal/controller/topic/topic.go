package topic

import (
	"context"

	v1 "focus-single/api/v1/topic"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	req.Type = consts.ContentTypeTopic
	out, err := content.GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	title := view.GetTitle(ctx, &model.ViewGetTitleInput{
		ContentType: req.Type,
		CategoryId:  req.CategoryId,
	})
	view.Render(ctx, model.View{
		ContentType: req.Type,
		Data:        out,
		Title:       title,
	})
	return nil, nil
}

func (c *controller) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	out, err := content.GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		view.Render404(ctx)
		return nil, nil
	}
	err = content.AddViewCount(ctx, req.Id, 1)
	view.Render(ctx, model.View{
		ContentType: consts.ContentTypeTopic,
		Data:        out,
		Title: view.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
			CurrentName: out.Content.Title,
		}),
		BreadCrumb: view.GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
			ContentId:   out.Content.Id,
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
		}),
	})
	return nil, nil
}
