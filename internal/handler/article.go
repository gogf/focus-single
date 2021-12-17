package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// 文章管理
	Article = handlerArticle{}
)

type handlerArticle struct{}

func (a *handlerArticle) Index(ctx context.Context, req *apiv1.ArticleIndexReq) (res *apiv1.ArticleIndexRes, err error) {
	req.Type = consts.ContentTypeArticle
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
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: req.Type,
				CategoryId:  req.CategoryId,
			}),
		})
	}
	return
}

func (a *handlerArticle) Detail(ctx context.Context, req *apiv1.ArticleDetailReq) (res *apiv1.ArticleDetailRes, err error) {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		if getDetailRes == nil {
			service.View.Render404(ctx)
			return nil, nil
		}
		if err = service.Content.AddViewCount(ctx, req.Id, 1); err != nil {
			return res, err
		}
		service.View.Render(ctx, model.View{
			ContentType: consts.ContentTypeArticle,
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
	}
	return
}
