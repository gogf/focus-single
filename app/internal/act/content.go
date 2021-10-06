package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 内容管理
	Content = contentAct{}
)

type contentAct struct{}

// @summary 展示创建内容页面
// @tags    前台-内容
// @produce html
// @router  /content/create [GET]
// @success 200 {string} html "页面HTML"
func (a *contentAct) Create(ctx context.Context, req *api.ContentCreateReq) (res *api.ContentCreateRes, err error) {
	service.View.Render(ctx, model.View{
		ContentType: req.Type,
	})
	return
}

// @summary 创建内容
// @description 客户端AJAX提交，客户端
// @tags    前台-内容
// @produce json
// @param   entity body internal.ContentDoCreateReq true "请求参数" required
// @router  /content/do-create [POST]
// @success 200 {object} model.ContentCreateOutput "请求结果"
func (a *contentAct) DoCreate(ctx context.Context, req *api.ContentDoCreateReq) (res *api.ContentDoCreateRes, err error) {
	out, err := service.Content.Create(ctx, req.ContentCreateInput)
	if err != nil {
		return nil, err
	}
	return &api.ContentDoCreateRes{ContentId: out.ContentId}, nil
}

// @summary 展示修改内容页面
// @tags    前台-内容
// @produce html
// @param   id query int true "问答ID"
// @router  /content/update [GET]
// @success 200 {string} html "页面HTML"
func (a *contentAct) Update(ctx context.Context, req *api.ContentUpdateReq) (res *api.ContentUpdateRes, err error) {
	if getDetailRes, err := service.Content.GetDetail(ctx, req.Id); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: getDetailRes.Content.Type,
			Data:        getDetailRes,
		})
	}
	return
}

// @summary 修改内容
// @tags    前台-内容
// @produce json
// @param   entity body internal.ContentDoUpdateReq true "请求参数" required
// @router  /content/do-update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentAct) DoUpdate(ctx context.Context, req *api.ContentDoUpdateReq) (res *api.ContentDoUpdateRes, err error) {
	err = service.Content.Update(ctx, req.ContentUpdateInput)
	return
}

// @summary 删除内容
// @tags    前台-内容
// @produce json
// @param   id formData int true "内容ID"
// @router  /content/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentAct) DoDelete(ctx context.Context, req *api.ContentDoDeleteReq) (res *api.ContentDoDeleteRes, err error) {
	err = service.Content.Delete(ctx, req.Id)
	return
}

// @summary 采纳回复
// @tags    前台-内容
// @produce json
// @param   entity body internal.ContentAdoptReplyReq true "请求参数" required
// @router  /content/adopt-reply [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentAct) AdoptReply(ctx context.Context, req *api.ContentAdoptReplyReq) (res *api.ContentAdoptReplyRes, err error) {
	err = service.Content.AdoptReply(ctx, req.Id, req.ReplyId)
	return
}
