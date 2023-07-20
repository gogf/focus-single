package interact

import (
	"context"

	"focus-single/api/interact/v1"
	"focus-single/internal/service"
)

func (c *ControllerV1) InteractCai(ctx context.Context, req *v1.InteractCaiReq) (res *v1.InteractCaiRes, err error) {
	err = service.Interact().Cai(ctx, req.Type, req.Id)
	return
}
