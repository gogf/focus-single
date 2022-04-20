package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

// Article 文章管理
var Article = cArticle{}

type cArticle struct{}

// Index article list
func (a *cArticle) Index(ctx context.Context, req *v1.ArticleIndexReq) (res *v1.ArticleIndexRes, err error) {
	req.Type = consts.ContentTypeArticle
	getListRes, err := content.GetList(ctx, content.GetListInput{
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
		ContentType: req.Type,
		Data:        getListRes,
		Title: view.GetTitle(ctx, &view.GetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		}),
	})
	return
}

// Detail .article details
func (a *cArticle) Detail(ctx context.Context, req *v1.ArticleDetailReq) (res *v1.ArticleDetailRes, err error) {
	getDetailRes, err := content.GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if getDetailRes == nil {
		view.Render404(ctx)
		return nil, nil
	}
	if err = content.AddViewCount(ctx, req.Id, 1); err != nil {
		return res, err
	}
	view.Render(ctx, view.View{
		ContentType: consts.ContentTypeArticle,
		Data:        getDetailRes,
		Title: view.GetTitle(ctx, &view.GetTitleInput{
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
			CurrentName: getDetailRes.Content.Title,
		}),
		BreadCrumb: view.GetBreadCrumb(ctx, &view.GetBreadCrumbInput{
			ContentId:   getDetailRes.Content.Id,
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
		}),
	})
	return
}
