package api

import (
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 主题管理
var Topic = topicApi{}

type topicApi struct{}

// @summary 展示主题首页
// @tags    前台-主题
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /topic [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Index(r *ghttp.Request) {
	var (
		req *internal.ContentGetListReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	req.Type = model.ContentTypeTopic
	if getListRes, err := service.Content.GetList(r.Context(), req.ContentGetListInput); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		title := service.View.GetTitle(r.Context(), &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		})
		service.View.Render(r, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       title,
		})
	}
}

// @summary 展示主题详情
// @tags    前台-主题
// @produce html
// @param   id path int false "主题ID"
// @router  /topic/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Detail(r *ghttp.Request) {
	var (
		data *internal.ContentDetailReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getDetailRes, err := service.Content.GetDetail(r.Context(), data.Id); err != nil {
		service.View.Render500(r)
	} else {
		if getDetailRes == nil {
			service.View.Render404(r)
		}
		service.Content.AddViewCount(r.Context(), data.Id, 1)
		service.View.Render(r, model.View{
			ContentType: model.ContentTypeTopic,
			Data:        getDetailRes,
			Title: service.View.GetTitle(r.Context(), &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			}),
			BreadCrumb: service.View.GetBreadCrumb(r.Context(), &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			}),
		})
	}
}
