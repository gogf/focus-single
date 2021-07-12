package api

import (
	"focus/app/api/internal"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
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
func (a *interactApi) Zan(r *ghttp.Request) {
	var (
		req *internal.InteractZanReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.Zan(r.Context(), req.Type, req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消赞
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) CancelZan(r *ghttp.Request) {
	var (
		req *internal.InteractCancelZanReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.CancelZan(r.Context(), req.Type, req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) Cai(r *ghttp.Request) {
	var (
		req *internal.InteractCaiReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.Cai(r.Context(), req.Type, req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) CancelCai(r *ghttp.Request) {
	var (
		req *internal.InteractCancelCaiReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.CancelCai(r.Context(), req.Type, req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}
