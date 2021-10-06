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

// @summary 赞
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactAct) Zan(ctx context.Context, req *api.InteractZanReq) (res *api.InteractZanRes, err error) {
	err = service.Interact.Zan(ctx, req.Type, req.Id)
	return
}

// @summary 取消赞
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactAct) CancelZan(ctx context.Context, req *api.InteractCancelZanReq) (res *api.InteractCancelZanRes, err error) {
	err = service.Interact.CancelZan(ctx, req.Type, req.Id)
	return
}

// @summary 踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactAct) Cai(ctx context.Context, req *api.InteractCaiReq) (res *api.InteractCaiRes, err error) {
	err = service.Interact.Cai(ctx, req.Type, req.Id)
	return
}

// @summary 取消踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactAct) CancelCai(ctx context.Context, req *api.InteractCancelCaiReq) (res *api.InteractCancelCaiRes, err error) {
	err = service.Interact.CancelCai(ctx, req.Type, req.Id)
	return
}
