package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
)

// 回复控制器
var Reply = replyApi{}

type replyApi struct{}

// @summary 创建回复
// @description 客户端AJAX提交，客户端
// @tags    前台-回复
// @produce json
// @param   entity body internal.ReplyDoCreateReq true "请求参数" required
// @router  /reply/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoCreate(ctx context.Context, req *internal.ReplyDoCreateReq) error {
	return service.Reply.Create(ctx, req.ReplyCreateInput)
}

// 获取回复列表
func (a *replyApi) Index(ctx context.Context, req *internal.ReplyGetListReq) (string, error) {
	if getListRes, err := service.Reply.GetList(ctx, req.ReplyGetListInput); err != nil {
		return "", err
	} else {
		request := g.RequestFromCtx(ctx)
		service.View.RenderTpl(ctx, "index/reply.html", model.View{Data: getListRes})
		tplContent := request.Response.BufferString()
		request.Response.ClearBuffer()
		return tplContent, nil
	}
}

// @summary 删除回复
// @tags    前台-回复
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoDelete(ctx context.Context, req *internal.ReplyDoDeleteReq) error {
	return service.Reply.Delete(ctx, req.Id)
}
