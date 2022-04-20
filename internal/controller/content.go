package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/service/content"
	"focus-single/internal/service/session"
	"focus-single/internal/service/view"
)

// Content 内容管理
var Content = cContent{}

type cContent struct{}

func (a *cContent) ShowCreate(ctx context.Context, req *v1.ContentShowCreateReq) (res *v1.ContentShowCreateRes, err error) {
	view.Render(ctx, view.View{
		ContentType: req.Type,
	})
	return
}

func (a *cContent) Create(ctx context.Context, req *v1.ContentCreateReq) (res *v1.ContentCreateRes, err error) {
	out, err := content.Create(ctx, content.CreateInput{
		CreateUpdateBase: content.CreateUpdateBase{
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
	return &v1.ContentCreateRes{ContentId: out.ContentId}, nil
}

func (a *cContent) ShowUpdate(ctx context.Context, req *v1.ContentShowUpdateReq) (res *v1.ContentShowUpdateRes, err error) {
	getDetailRes, err := content.GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	view.Render(ctx, view.View{
		ContentType: getDetailRes.Content.Type,
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

func (a *cContent) Update(ctx context.Context, req *v1.ContentUpdateReq) (res *v1.ContentUpdateRes, err error) {
	err = content.Update(ctx, content.UpdateInput{
		Id: req.Id,
		CreateUpdateBase: content.CreateUpdateBase{
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
	err = content.Delete(ctx, req.Id)
	return
}
