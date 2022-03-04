package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

// 搜索管理
var Search = cSearch{}

type cSearch struct{}

func (a *cSearch) Index(ctx context.Context, req *v1.SearchIndexReq) (res *v1.SearchIndexRes, err error) {
	if searchRes, err := service.Content().Search(ctx, model.ContentSearchInput{
		Key:        req.Key,
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		service.View().Render(ctx, model.View{
			Data:  searchRes,
			Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{}),
		})
		return nil, nil
	}
}
