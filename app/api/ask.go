package api

import (
	"focus/app/api/internal"
	"focus/app/cnt"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 问答管理
var Ask = askApi{}

type askApi struct{}

// @summary 展示问答首页
// @tags    前台-问答
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /ask [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Index(r *ghttp.Request) {
	var (
		req *internal.ContentGetListReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	req.Type = cnt.ContentTypeAsk
	if getListRes, err := service.Content.GetList(r.Context(), req.ContentGetListInput); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title: service.View.GetTitle(r.Context(), &model.ViewGetTitleInput{
				ContentType: req.Type,
				CategoryId:  req.CategoryId,
			}),
		})
	}
}

// @summary 展示问答详情
// @tags    前台-问答
// @produce html
// @param   id path int false "问答ID"
// @router  /ask/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Detail(r *ghttp.Request) {
	var (
		req *internal.ContentDetailReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getDetailRes, err := service.Content.GetDetail(r.Context(), req.Id); err != nil {
		service.View.Render500(r)
	} else {
		if getDetailRes == nil {
			service.View.Render404(r)
		}
		if err := service.Content.AddViewCount(r.Context(), req.Id, 1); err != nil {
			service.View.Render500(r, model.View{
				Error: err.Error(),
			})
		}
		var (
			title = service.View.GetTitle(r.Context(), &model.ViewGetTitleInput{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			})
			breadCrumb = service.View.GetBreadCrumb(r.Context(), &model.ViewGetBreadCrumbInput{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			})
		)

		service.View.Render(r, model.View{
			ContentType: cnt.ContentTypeAsk,
			Data:        getDetailRes,
			Title:       title,
			BreadCrumb:  breadCrumb,
		})
	}
}
