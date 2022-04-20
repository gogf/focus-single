package controller

import (
	"context"

	"focus-single/internal/service/reply"
	"focus-single/internal/service/session"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/frame/g"

	"focus-single/api/v1"
)

// Reply 回复控制器
var Reply = cReply{}

type cReply struct{}

func (a *cReply) GetListContent(ctx context.Context, req *v1.ReplyGetListContentReq) (res *v1.ReplyGetListContentRes, err error) {
	out, err := reply.GetList(ctx, reply.GetListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
	})
	if err != nil {
		return nil, err
	}
	request := g.RequestFromCtx(ctx)
	view.RenderTpl(ctx, "index/reply.html", view.View{Data: out})
	tplContent := request.Response.BufferString()
	request.Response.ClearBuffer()
	return &v1.ReplyGetListContentRes{Content: tplContent}, nil
}

func (a *cReply) Create(ctx context.Context, req *v1.ReplyCreateReq) (res *v1.ReplyCreateRes, err error) {
	err = reply.Create(ctx, reply.CreateInput{
		Title:      req.Title,
		ParentId:   req.ParentId,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		Content:    req.Content,
		UserId:     session.GetUser(ctx).Id,
	})
	return
}

func (a *cReply) Delete(ctx context.Context, req *v1.ReplyDeleteReq) (res *v1.ReplyDeleteRes, err error) {
	err = reply.Delete(ctx, req.Id)
	return
}
