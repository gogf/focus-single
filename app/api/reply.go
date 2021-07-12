package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 回复控制器
var Reply = replyApi{}

type replyApi struct{}

// @summary 创建回复
// @description 客户端AJAX提交，客户端
// @tags    前台-回复
// @produce json
// @param   entity body model.ReplyDoCreateReq true "请求参数" required
// @router  /reply/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoCreate(r *ghttp.Request) {
	var (
		req *model.ReplyDoCreateReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Reply.Create(r.Context(), req.ReplyCreateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// 回复列表
func (a *replyApi) Index(r *ghttp.Request) {
	var (
		req *model.ReplyGetListReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if getListRes, err := service.Reply.GetList(r.Context(), req.ReplyGetListInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		service.View.RenderTpl(r, "index/reply.html", model.View{Data: getListRes})
	}
}

// @summary 删除回复
// @tags    前台-回复
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoDelete(r *ghttp.Request) {
	var (
		req *model.ReplyDoDeleteReq
	)
	if err := r.ParseForm(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Reply.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}
