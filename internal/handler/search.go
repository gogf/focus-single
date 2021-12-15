package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// 搜索管理
	Search = handlerSearch{}
)

type handlerSearch struct{}

func (a *handlerSearch) Index(ctx context.Context, req *apiv1.ContentSearchReq) (res *apiv1.ContentSearchRes, err error) {
	if searchRes, err := service.Content.Search(ctx, model.ContentSearchInput{
		Key:        req.Key,
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			Data:  searchRes,
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{}),
		})
		return nil, nil
	}
}
