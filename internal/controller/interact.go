package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/service/interact"
)

// 赞踩控制器
var Interact = cInteract{}

type cInteract struct{}

func (a *cInteract) Zan(ctx context.Context, req *v1.InteractZanReq) (res *v1.InteractZanRes, err error) {
	err = interact.Zan(ctx, req.Type, req.Id)
	return
}

func (a *cInteract) CancelZan(ctx context.Context, req *v1.InteractCancelZanReq) (res *v1.InteractCancelZanRes, err error) {
	err = interact.CancelZan(ctx, req.Type, req.Id)
	return
}

func (a *cInteract) Cai(ctx context.Context, req *v1.InteractCaiReq) (res *v1.InteractCaiRes, err error) {
	err = interact.Cai(ctx, req.Type, req.Id)
	return
}

func (a *cInteract) CancelCai(ctx context.Context, req *v1.InteractCancelCaiReq) (res *v1.InteractCancelCaiRes, err error) {
	err = interact.CancelCai(ctx, req.Type, req.Id)
	return
}
