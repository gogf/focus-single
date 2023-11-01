package reply

import (
	"context"

	"focus-single/api/reply/v1"
	"focus-single/internal/service"
)

func (c *ControllerV1) ReplyDelete(ctx context.Context, req *v1.ReplyDeleteReq) (res *v1.ReplyDeleteRes, err error) {
	err = service.Reply().Delete(ctx, req.Id)
	return
}
