package controller

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

// Ask 问答管理
var Ask = cAak{}

type cAak struct{}

func (a *cAak) Index(ctx context.Context, req *apiv1.AskIndexReq) (res *apiv1.AskIndexRes, err error) {
	req.Type = consts.ContentTypeAsk
	if getListRes, err := service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		service.View().Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: req.Type,
				CategoryId:  req.CategoryId,
			}),
		})
	}
	return
}

func (a *cAak) Detail(ctx context.Context, req *apiv1.AskDetailReq) (res *apiv1.AskDetailRes, err error) {
	if getDetailRes, err := service.Content().GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		if getDetailRes == nil {
			service.View().Render404(ctx)
			return nil, nil
		}
		if err = service.Content().AddViewCount(ctx, req.Id, 1); err != nil {
			return nil, err
		}
		var (
			title = service.View().GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			})
			breadCrumb = service.View().GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			})
		)
		service.View().Render(ctx, model.View{
			ContentType: consts.ContentTypeAsk,
			Data:        getDetailRes,
			Title:       title,
			BreadCrumb:  breadCrumb,
		})
	}
	return
}
