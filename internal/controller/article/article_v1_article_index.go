package article

import (
	"context"

	"focus-single/api/article/v1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ArticleIndex(ctx context.Context, req *v1.ArticleIndexReq) (res *v1.ArticleIndexRes, err error) {
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
