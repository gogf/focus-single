package search

import (
	"context"

	v1 "focus-single/api/v1/search"
	"focus-single/internal/model"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	if searchRes, err := content.Search(ctx, model.ContentSearchInput{
		Key:        req.Key,
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		view.Render(ctx, model.View{
			Data:  searchRes,
			Title: view.GetTitle(ctx, &model.ViewGetTitleInput{}),
		})
		return nil, nil
	}
}
