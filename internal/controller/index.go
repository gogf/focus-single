package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

// 首页接口
var Index = cIndex{}

type cIndex struct{}

func (a *cIndex) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	if getListRes, err := content.GetList(ctx, content.GetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		view.Render(ctx, view.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       "首页",
		})
	}
	return
}
