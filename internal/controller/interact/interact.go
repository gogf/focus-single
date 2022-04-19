package interact

import (
	"context"

	v1 "focus-single/api/v1/interact"
	"focus-single/internal/service/interact"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Zan(ctx context.Context, req *v1.ZanReq) (res *v1.ZanRes, err error) {
	err = interact.Zan(ctx, req.Type, req.Id)
	return
}

func (c *controller) CancelZan(ctx context.Context, req *v1.CancelZanReq) (res *v1.CancelZanRes, err error) {
	err = interact.CancelZan(ctx, req.Type, req.Id)
	return
}

func (c *controller) Cai(ctx context.Context, req *v1.CaiReq) (res *v1.CaiRes, err error) {
	err = interact.Cai(ctx, req.Type, req.Id)
	return
}

func (c *controller) CancelCai(ctx context.Context, req *v1.CancelCaiReq) (res *v1.CancelCaiRes, err error) {
	err = interact.CancelCai(ctx, req.Type, req.Id)
	return
}
