package controller

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

// Article 文章管理
var Article = cArticle{}

type cArticle struct{}

// Index article list
func (a *cArticle) Index(ctx context.Context, req *apiv1.ArticleIndexReq) (res *apiv1.ArticleIndexRes, err error) {
	req.Type = consts.ContentTypeArticle
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
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
		Data:        getListRes,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		}),
	})
	return
}

// Detail .article details
func (a *cArticle) Detail(ctx context.Context, req *apiv1.ArticleDetailReq) (res *apiv1.ArticleDetailRes, err error) {
	getDetailRes, err := service.Content().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if getDetailRes == nil {
		service.View().Render404(ctx)
		return nil, nil
	}
	if err = service.Content().AddViewCount(ctx, req.Id, 1); err != nil {
		return res, err
	}
	service.View().Render(ctx, model.View{
		ContentType: consts.ContentTypeArticle,
		Data:        getDetailRes,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
			CurrentName: getDetailRes.Content.Title,
		}),
		BreadCrumb: service.View().GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
			ContentId:   getDetailRes.Content.Id,
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
		}),
	})
	return
}
