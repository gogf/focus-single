package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// Search 搜索管理
	Search = handlerSearch{}
)

type handlerSearch struct{}

func (a *handlerSearch) Index(ctx context.Context, req *apiv1.SearchIndexReq) (res *apiv1.SearchIndexRes, err error) {
	searchRes, err := service.Content().Search(ctx, model.ContentSearchInput{
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
	service.View.Render(ctx, model.View{
		Data:  searchRes,
		Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{}),
	})
	return nil, nil
}
