package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 内容管理
	Content = contentAct{}
)

type contentAct struct{}

func (a *contentAct) Create(ctx context.Context, req *api.ContentCreateReq) (res *api.ContentCreateRes, err error) {
	service.View.Render(ctx, model.View{
		ContentType: req.Type,
	})
	return
}

func (a *contentAct) DoCreate(ctx context.Context, req *api.ContentDoCreateReq) (res *api.ContentDoCreateRes, err error) {
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
	return &api.ContentDoCreateRes{ContentId: out.ContentId}, nil
}

func (a *contentAct) Update(ctx context.Context, req *api.ContentUpdateReq) (res *api.ContentUpdateRes, err error) {
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

func (a *contentAct) DoUpdate(ctx context.Context, req *api.ContentDoUpdateReq) (res *api.ContentDoUpdateRes, err error) {
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

func (a *contentAct) DoDelete(ctx context.Context, req *api.ContentDoDeleteReq) (res *api.ContentDoDeleteRes, err error) {
	err = service.Content.Delete(ctx, req.Id)
	return
}

func (a *contentAct) AdoptReply(ctx context.Context, req *api.ContentAdoptReplyReq) (res *api.ContentAdoptReplyRes, err error) {
	err = service.Content.AdoptReply(ctx, req.Id, req.ReplyId)
	return
}
