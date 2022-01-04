package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// Index 首页接口
	Index = handlerIndex{}
)

type handlerIndex struct{}

func (a *handlerIndex) Index(ctx context.Context, req *apiv1.IndexReq) (res *apiv1.IndexRes, err error) {
	getListRes, err := service.Content().GetList(ctx, model.ContentGetListInput{
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
		ContentType: req.Type,
		Data:        getListRes,
		Title:       "首页",
	})
	return
}
