package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/service"
)

var (
	// Interact 赞踩控制器
	Interact = handlerInteract{}
)

type handlerInteract struct{}

func (a *handlerInteract) Zan(ctx context.Context, req *apiv1.InteractZanReq) (res *apiv1.InteractZanRes, err error) {
	err = service.Interact.Zan(ctx, req.Type, req.Id)
	return
}

func (a *handlerInteract) CancelZan(ctx context.Context, req *apiv1.InteractCancelZanReq) (res *apiv1.InteractCancelZanRes, err error) {
	err = service.Interact.CancelZan(ctx, req.Type, req.Id)
	return
}

func (a *handlerInteract) Cai(ctx context.Context, req *apiv1.InteractCaiReq) (res *apiv1.InteractCaiRes, err error) {
	err = service.Interact.Cai(ctx, req.Type, req.Id)
	return
}

func (a *handlerInteract) CancelCai(ctx context.Context, req *apiv1.InteractCancelCaiReq) (res *apiv1.InteractCancelCaiRes, err error) {
	err = service.Interact.CancelCai(ctx, req.Type, req.Id)
	return
}
