package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// 回复控制器
	Reply = handlerReply{}
)

type handlerReply struct{}

func (a *handlerReply) Index(ctx context.Context, req *apiv1.ReplyGetListReq) (res *apiv1.ReplyGetListRes, err error) {
	if getListRes, err := service.Reply.GetList(ctx, model.ReplyGetListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
	}); err != nil {
		return nil, err
	} else {
		request := g.RequestFromCtx(ctx)
		service.View.RenderTpl(ctx, "index/reply.html", model.View{Data: getListRes})
		tplContent := request.Response.BufferString()
		request.Response.ClearBuffer()
		return &apiv1.ReplyGetListRes{Content: tplContent}, nil
	}
}

func (a *handlerReply) DoCreate(ctx context.Context, req *apiv1.ReplyDoCreateReq) (res *apiv1.ReplyDoCreateRes, err error) {
	err = service.Reply.Create(ctx, model.ReplyCreateInput{
		Title:      req.Title,
		ParentId:   req.ParentId,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		Content:    req.Content,
		UserId:     service.Session.GetUser(ctx).Id,
	})
	return
}

func (a *handlerReply) DoDelete(ctx context.Context, req *apiv1.ReplyDoDeleteReq) (res *apiv1.ReplyDoDeleteRes, err error) {
	err = service.Reply.Delete(ctx, req.Id)
	return
}
