package reply

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"focus-single/api/reply/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ReplyGetListContent(ctx context.Context, req *v1.ReplyGetListContentReq) (res *v1.ReplyGetListContentRes, err error) {
	out, err := service.Reply().GetList(ctx, model.ReplyGetListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
	})
	if err != nil {
		return nil, err
	}
	request := g.RequestFromCtx(ctx)
	service.View().RenderTpl(ctx, "index/reply.html", model.View{Data: out})
	tplContent := request.Response.BufferString()
	request.Response.ClearBuffer()
	return &v1.ReplyGetListContentRes{Content: tplContent}, nil
}
