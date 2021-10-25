package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 搜索管理
	Search = searchAct{}
)

type searchAct struct{}

func (a *searchAct) Index(ctx context.Context, req *api.ContentSearchReq) (res *api.ContentSearchRes, err error) {
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
