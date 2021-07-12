package api

import (
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 首页接口
var Index = indexApi{}

type indexApi struct{}

// @summary 展示站点首页
// @tags    前台-首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func (a *indexApi) Index(r *ghttp.Request) {
	var (
		req *internal.ContentGetListReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getListRes, err := service.Content.GetList(r.Context(), req.ContentGetListInput); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       "首页",
		})
	}
}
