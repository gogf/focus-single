package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
	"github.com/gogf/gf/frame/g"
)

var (
	// 回复控制器
	Reply = replyAct{}
)

type replyAct struct{}

// @summary 创建回复
// @description 客户端AJAX提交，客户端
// @tags    前台-回复
// @produce json
// @param   entity body internal.ReplyDoCreateReq true "请求参数" required
// @router  /reply/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyAct) DoCreate(ctx context.Context, req *api.ReplyDoCreateReq) (res *api.ReplyDoCreateRes, err error) {
	err = service.Reply.Create(ctx, req.ReplyCreateInput)
	return
}

// 获取回复列表
func (a *replyAct) Index(ctx context.Context, req *api.ReplyGetListReq) (res *api.ReplyGetListRes, err error) {
	if getListRes, err := service.Reply.GetList(ctx, req.ReplyGetListInput); err != nil {
		return nil, err
	} else {
		request := g.RequestFromCtx(ctx)
		service.View.RenderTpl(ctx, "index/reply.html", model.View{Data: getListRes})
		tplContent := request.Response.BufferString()
		request.Response.ClearBuffer()
		return &api.ReplyGetListRes{Content: tplContent}, nil
	}
}

// @summary 删除回复
// @tags    前台-回复
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyAct) DoDelete(ctx context.Context, req *api.ReplyDoDeleteReq) (res *api.ReplyDoDeleteRes, err error) {
	err = service.Reply.Delete(ctx, req.Id)
	return
}
