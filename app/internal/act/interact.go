package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/service"
)

var (
	// 赞踩控制器
	Interact = interactAct{}
)

type interactAct struct{}

func (a *interactAct) Zan(ctx context.Context, req *api.InteractZanReq) (res *api.InteractZanRes, err error) {
	err = service.Interact.Zan(ctx, req.Type, req.Id)
	return
}

func (a *interactAct) CancelZan(ctx context.Context, req *api.InteractCancelZanReq) (res *api.InteractCancelZanRes, err error) {
	err = service.Interact.CancelZan(ctx, req.Type, req.Id)
	return
}

func (a *interactAct) Cai(ctx context.Context, req *api.InteractCaiReq) (res *api.InteractCaiRes, err error) {
	err = service.Interact.Cai(ctx, req.Type, req.Id)
	return
}

func (a *interactAct) CancelCai(ctx context.Context, req *api.InteractCancelCaiReq) (res *api.InteractCancelCaiRes, err error) {
	err = service.Interact.CancelCai(ctx, req.Type, req.Id)
	return
}
