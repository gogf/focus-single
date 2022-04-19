package reply

import (
	"context"

	v1 "focus-single/api/v1/reply"
	"focus-single/internal/model"
	"focus-single/internal/service/reply"
	"focus-single/internal/service/session"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/frame/g"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) GetListContent(ctx context.Context, req *v1.GetListContentReq) (res *v1.GetListContentRes, err error) {
	out, err := reply.GetList(ctx, model.ReplyGetListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
	})
	if err != nil {
		return nil, err
	}
	request := g.RequestFromCtx(ctx)
	view.RenderTpl(ctx, "index/reply.html", model.View{Data: out})
	tplContent := request.Response.BufferString()
	request.Response.ClearBuffer()
	return &v1.GetListContentRes{Content: tplContent}, nil
}

func (c *controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = reply.Create(ctx, model.ReplyCreateInput{
		Title:      req.Title,
		ParentId:   req.ParentId,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		Content:    req.Content,
		UserId:     session.GetUser(ctx).Id,
	})
	return
}

func (c *controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = reply.Delete(ctx, req.Id)
	return
}
