package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 首页接口
	Index = indexAct{}
)

type indexAct struct{}

func (a *indexAct) Index(ctx context.Context, req *api.ContentGetListReq) (res *api.ContentGetListRes, err error) {
	if getListRes, err := service.Content.GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
		UserId:     service.Session.GetUser(ctx).Id,
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
