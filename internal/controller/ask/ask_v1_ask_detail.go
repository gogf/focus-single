package ask

import (
	"context"

	"focus-single/api/ask/v1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) AskDetail(ctx context.Context, req *v1.AskDetailReq) (res *v1.AskDetailRes, err error) {
	getDetailRes, err := service.Content().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
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
	return
}
