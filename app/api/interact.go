package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/service"
)

// 赞踩控制器
var Interact = interactApi{}

type interactApi struct{}

// @summary 赞
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) Zan(ctx context.Context, req *internal.InteractZanReq) error {
	return service.Interact.Zan(ctx, req.Type, req.Id)
}

// @summary 取消赞
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) CancelZan(ctx context.Context, req *internal.InteractCancelZanReq) error {
	return service.Interact.CancelZan(ctx, req.Type, req.Id)
}

// @summary 踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) Cai(ctx context.Context, req *internal.InteractCaiReq) error {
	return service.Interact.Cai(ctx, req.Type, req.Id)
}

// @summary 取消踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) CancelCai(ctx context.Context, req *internal.InteractCancelCaiReq) error {
	return service.Interact.CancelCai(ctx, req.Type, req.Id)
}
