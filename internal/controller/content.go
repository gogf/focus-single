package controller

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

// Content 内容管理
var Content = cContent{}

type cContent struct{}

func (a *cContent) ShowCreate(ctx context.Context, req *apiv1.ContentShowCreateReq) (res *apiv1.ContentShowCreateRes, err error) {
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
	})
	return
}

func (a *cContent) Create(ctx context.Context, req *apiv1.ContentCreateReq) (res *apiv1.ContentCreateRes, err error) {
	out, err := service.Content().Create(ctx, model.ContentCreateInput{
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
		UserId: service.Session().GetUser(ctx).Id,
	})
	if err != nil {
		return nil, err
	}
	return &apiv1.ContentCreateRes{ContentId: out.ContentId}, nil
}

func (a *cContent) ShowUpdate(ctx context.Context, req *apiv1.ContentShowUpdateReq) (res *apiv1.ContentShowUpdateRes, err error) {
	if getDetailRes, err := service.Content().GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		service.View().Render(ctx, model.View{
			ContentType: getDetailRes.Content.Type,
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
	}
	return
}

func (a *cContent) Update(ctx context.Context, req *apiv1.ContentUpdateReq) (res *apiv1.ContentUpdateRes, err error) {
	err = service.Content().Update(ctx, model.ContentUpdateInput{
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

func (a *cContent) Delete(ctx context.Context, req *apiv1.ContentDeleteReq) (res *apiv1.ContentDeleteRes, err error) {
	err = service.Content().Delete(ctx, req.Id)
	return
}
