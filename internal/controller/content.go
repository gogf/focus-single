package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

// Content 内容管理
var Content = cContent{}

type cContent struct{}

func (a *cContent) ShowCreate(ctx context.Context, req *v1.ContentShowCreateReq) (res *v1.ContentShowCreateRes, err error) {
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
	})
	return
}

func (a *cContent) Create(ctx context.Context, req *v1.ContentCreateReq) (res *v1.ContentCreateRes, err error) {
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
	return &v1.ContentCreateRes{ContentId: out.ContentId}, nil
}

func (a *cContent) ShowUpdate(ctx context.Context, req *v1.ContentShowUpdateReq) (res *v1.ContentShowUpdateRes, err error) {
	getDetailRes, err := service.Content().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
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
	return
}

func (a *cContent) Update(ctx context.Context, req *v1.ContentUpdateReq) (res *v1.ContentUpdateRes, err error) {
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

func (a *cContent) Delete(ctx context.Context, req *v1.ContentDeleteReq) (res *v1.ContentDeleteRes, err error) {
	err = service.Content().Delete(ctx, req.Id)
	return
}
