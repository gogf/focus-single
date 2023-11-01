package interact

import (
	"context"

	"focus-single/api/interact/v1"
	"focus-single/internal/service"
)

func (c *ControllerV1) InteractCancelZan(ctx context.Context, req *v1.InteractCancelZanReq) (res *v1.InteractCancelZanRes, err error) {
	err = service.Interact().CancelZan(ctx, req.Type, req.Id)
	return
}
