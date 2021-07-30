package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
)

// 内容管理
var Content = contentApi{}

type contentApi struct{}

// @summary 展示创建内容页面
// @tags    前台-内容
// @produce html
// @router  /content/create [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Create(ctx context.Context, req *internal.ContentCreateReq) {
	service.View.Render(ctx, model.View{
		ContentType: req.Type,
	})
}

// @summary 创建内容
// @description 客户端AJAX提交，客户端
// @tags    前台-内容
// @produce json
// @param   entity body internal.ContentDoCreateReq true "请求参数" required
// @router  /content/do-create [POST]
// @success 200 {object} model.ContentCreateOutput "请求结果"
func (a *contentApi) DoCreate(ctx context.Context, req *internal.ContentDoCreateReq) (output model.ContentCreateOutput, err error) {
	return service.Content.Create(ctx, req.ContentCreateInput)
}

// @summary 展示修改内容页面
// @tags    前台-内容
// @produce html
// @param   id query int true "问答ID"
// @router  /content/update [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Update(ctx context.Context, req *internal.ContentUpdateReq) error {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: getDetailRes.Content.Type,
			Data:        getDetailRes,
		})
	}
	return nil
}

// @summary 修改内容
// @tags    前台-内容
// @produce json
// @param   entity body internal.ContentDoUpdateReq true "请求参数" required
// @router  /content/do-update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoUpdate(ctx context.Context, req *internal.ContentDoUpdateReq) error {
	return service.Content.Update(ctx, req.ContentUpdateInput)
}

// @summary 删除内容
// @tags    前台-内容
// @produce json
// @param   id formData int true "内容ID"
// @router  /content/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoDelete(ctx context.Context, req *internal.ContentDoDeleteReq) error {
	return service.Content.Delete(ctx, req.Id)
}

// @summary 采纳回复
// @tags    前台-内容
// @produce json
// @param   entity body internal.ContentAdoptReplyReq true "请求参数" required
// @router  /content/adopt-reply [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) AdoptReply(ctx context.Context, req *internal.ContentAdoptReplyReq) error {
	return service.Content.AdoptReply(ctx, req.Id, req.ReplyId)
}
