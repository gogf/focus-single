package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// 主题管理
	Topic = handlerTopic{}
)

type handlerTopic struct{}

func (a *handlerTopic) Index(ctx context.Context, req *apiv1.ContentGetListReq) (res *apiv1.ContentGetListRes, err error) {
	req.Type = consts.ContentTypeTopic
	if getListRes, err := service.Content.GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		title := service.View.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.ContentListCommonReq.CategoryId,
		})
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       title,
		})
		return nil, nil
	}
}

func (a *handlerTopic) Detail(ctx context.Context, req *apiv1.ContentDetailReq) (res *apiv1.ContentDetailRes, err error) {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		if getDetailRes == nil {
			service.View.Render404(ctx)
			return nil, nil
		}
		err = service.Content.AddViewCount(ctx, req.Id, 1)
		service.View.Render(ctx, model.View{
			ContentType: consts.ContentTypeTopic,
			Data:        getDetailRes,
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			}),
			BreadCrumb: service.View.GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			}),
		})
		return nil, nil
	}
}
