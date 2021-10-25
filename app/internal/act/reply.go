package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// 回复控制器
	Reply = replyAct{}
)

type replyAct struct{}

func (a *replyAct) Index(ctx context.Context, req *api.ReplyGetListReq) (res *api.ReplyGetListRes, err error) {
	if getListRes, err := service.Reply.GetList(ctx, model.ReplyGetListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		UserId:     service.Session.GetUser(ctx).Id,
	}); err != nil {
		return nil, err
	} else {
		request := g.RequestFromCtx(ctx)
		service.View.RenderTpl(ctx, "index/reply.html", model.View{Data: getListRes})
		tplContent := request.Response.BufferString()
		request.Response.ClearBuffer()
		return &api.ReplyGetListRes{Content: tplContent}, nil
	}
}

func (a *replyAct) DoCreate(ctx context.Context, req *api.ReplyDoCreateReq) (res *api.ReplyDoCreateRes, err error) {
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

func (a *replyAct) DoDelete(ctx context.Context, req *api.ReplyDoDeleteReq) (res *api.ReplyDoDeleteRes, err error) {
	err = service.Reply.Delete(ctx, req.Id)
	return
}
