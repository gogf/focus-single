package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// 首页接口
	Index = handlerIndex{}
)

type handlerIndex struct{}

func (a *handlerIndex) Index(ctx context.Context, req *apiv1.ContentGetListReq) (res *apiv1.ContentGetListRes, err error) {
	if getListRes, err := service.Content.GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       "首页",
		})
	}
	return
}
