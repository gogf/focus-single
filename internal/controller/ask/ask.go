package ask

import (
	"context"

	v1 "focus-single/api/v1/ask"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service/content"
	"focus-single/internal/service/view"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	req.Type = consts.ContentTypeAsk
	if getListRes, err := content.GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		view.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title: view.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: req.Type,
				CategoryId:  req.CategoryId,
			}),
		})
	}
	return
}

func (c *controller) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	if getDetailRes, err := content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		if getDetailRes == nil {
			view.Render404(ctx)
			return nil, nil
		}
		if err = content.AddViewCount(ctx, req.Id, 1); err != nil {
			return nil, err
		}
		var (
			title = view.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			})
			breadCrumb = view.GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			})
		)
		view.Render(ctx, model.View{
			ContentType: consts.ContentTypeAsk,
			Data:        getDetailRes,
			Title:       title,
			BreadCrumb:  breadCrumb,
		})
	}
	return
}
