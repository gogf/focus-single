package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

var (
	// 内容管理
	Content = handlerContent{}
)

type handlerContent struct{}

func (a *handlerContent) Create(ctx context.Context, req *apiv1.ContentCreateReq) (res *apiv1.ContentCreateRes, err error) {
	service.View.Render(ctx, model.View{
		ContentType: req.Type,
	})
	return
}

func (a *handlerContent) DoCreate(ctx context.Context, req *apiv1.ContentDoCreateReq) (res *apiv1.ContentDoCreateRes, err error) {
	out, err := service.Content.Create(ctx, model.ContentCreateInput{
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
		UserId: service.Session.GetUser(ctx).Id,
	})
	if err != nil {
		return nil, err
	}
	return &apiv1.ContentDoCreateRes{ContentId: out.ContentId}, nil
}

func (a *handlerContent) Update(ctx context.Context, req *apiv1.ContentUpdateReq) (res *apiv1.ContentUpdateRes, err error) {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: getDetailRes.Content.Type,
			Data:        getDetailRes,
		})
	}
	return
}

func (a *handlerContent) DoUpdate(ctx context.Context, req *apiv1.ContentDoUpdateReq) (res *apiv1.ContentDoUpdateRes, err error) {
	err = service.Content.Update(ctx, model.ContentUpdateInput{
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

func (a *handlerContent) DoDelete(ctx context.Context, req *apiv1.ContentDoDeleteReq) (res *apiv1.ContentDoDeleteRes, err error) {
	err = service.Content.Delete(ctx, req.Id)
	return
}

func (a *handlerContent) AdoptReply(ctx context.Context, req *apiv1.ContentAdoptReplyReq) (res *apiv1.ContentAdoptReplyRes, err error) {
	err = service.Content.AdoptReply(ctx, req.Id, req.ReplyId)
	return
}
