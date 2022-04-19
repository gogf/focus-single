package content

import (
	"context"

	v1 "focus-single/api/v1/content"
	"focus-single/internal/model"
	"focus-single/internal/service/content"
	"focus-single/internal/service/session"
	"focus-single/internal/service/view"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) ShowCreate(ctx context.Context, req *v1.ShowCreateReq) (res *v1.ShowCreateRes, err error) {
	view.Render(ctx, model.View{
		ContentType: req.Type,
	})
	return
}

func (c *controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	out, err := content.Create(ctx, model.ContentCreateInput{
		ContentCreateUpdateBase: model.ContentCreateUpdateBase{
			Type:       req.Type,
			CategoryId: req.CategoryId,
			Title:      req.Title,
			Content:    req.Content,
			Brief:      req.Brief,
			Thumb:      req.Thumb,
			Tags:       req.Tags,
			Referer:    req.Referer,
		},
		UserId: session.GetUser(ctx).Id,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{ContentId: out.ContentId}, nil
}

func (c *controller) ShowUpdate(ctx context.Context, req *v1.ShowUpdateReq) (res *v1.ShowUpdateRes, err error) {
	if getDetailRes, err := content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		view.Render(ctx, model.View{
			ContentType: getDetailRes.Content.Type,
			Data:        getDetailRes,
			Title: view.GetTitle(ctx, &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			}),
			BreadCrumb: view.GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			}),
		})
	}
	return
}

func (c *controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = content.Update(ctx, model.ContentUpdateInput{
		Id: req.Id,
		ContentCreateUpdateBase: model.ContentCreateUpdateBase{
			Type:       req.Type,
			CategoryId: req.CategoryId,
			Title:      req.Title,
			Content:    req.Content,
			Brief:      req.Brief,
			Thumb:      req.Thumb,
			Tags:       req.Tags,
			Referer:    req.Referer,
		},
	})
	return
}

func (c *controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = content.Delete(ctx, req.Id)
	return
}
