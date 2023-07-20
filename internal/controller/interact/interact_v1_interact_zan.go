package interact

import (
	"context"

	"focus-single/api/interact/v1"
	"focus-single/internal/service"
)

func (c *ControllerV1) InteractZan(ctx context.Context, req *v1.InteractZanReq) (res *v1.InteractZanRes, err error) {
	err = service.Interact().Zan(ctx, req.Type, req.Id)
	return
}
