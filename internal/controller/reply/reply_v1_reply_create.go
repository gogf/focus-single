package reply

import (
	"context"

	"focus-single/api/reply/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ReplyCreate(ctx context.Context, req *v1.ReplyCreateReq) (res *v1.ReplyCreateRes, err error) {
	err = service.Reply().Create(ctx, model.ReplyCreateInput{
		Title:      req.Title,
		ParentId:   req.ParentId,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		Content:    req.Content,
		UserId:     service.Session().GetUser(ctx).Id,
	})
	return
}
