package interact

import (
	"context"

	"focus-single/api/interact/v1"
	"focus-single/internal/service"
)

func (c *ControllerV1) InteractCancelCai(ctx context.Context, req *v1.InteractCancelCaiReq) (res *v1.InteractCancelCaiRes, err error) {
	err = service.Interact().CancelCai(ctx, req.Type, req.Id)
	return
}
