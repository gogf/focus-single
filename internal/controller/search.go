package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

// 搜索管理
var Search = cSearch{}

type cSearch struct{}

func (a *cSearch) Index(ctx context.Context, req *v1.SearchIndexReq) (res *v1.SearchIndexRes, err error) {
	out, err := content.Search(ctx, content.SearchInput{
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
	view.Render(ctx, view.View{
		Data:  out,
		Title: view.GetTitle(ctx, &view.GetTitleInput{}),
	})
	return nil, nil
}
