package search

import (
	"context"

	"focus-single/api/search/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) SearchIndex(ctx context.Context, req *v1.SearchIndexReq) (res *v1.SearchIndexRes, err error) {
	out, err := service.Content().Search(ctx, model.ContentSearchInput{
		Key:        req.Key,
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	service.View().Render(ctx, model.View{
		Data:  out,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{}),
	})
	return nil, nil
}
