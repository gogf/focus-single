package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/service"
)

// 赞踩控制器
var Interact = cInteract{}

type cInteract struct{}

func (a *cInteract) Zan(ctx context.Context, req *v1.InteractZanReq) (res *v1.InteractZanRes, err error) {
	err = service.Interact().Zan(ctx, req.Type, req.Id)
	return
}

func (a *cInteract) CancelZan(ctx context.Context, req *v1.InteractCancelZanReq) (res *v1.InteractCancelZanRes, err error) {
	err = service.Interact().CancelZan(ctx, req.Type, req.Id)
	return
}

func (a *cInteract) Cai(ctx context.Context, req *v1.InteractCaiReq) (res *v1.InteractCaiRes, err error) {
	err = service.Interact().Cai(ctx, req.Type, req.Id)
	return
}

func (a *cInteract) CancelCai(ctx context.Context, req *v1.InteractCancelCaiReq) (res *v1.InteractCancelCaiRes, err error) {
	err = service.Interact().CancelCai(ctx, req.Type, req.Id)
	return
}
