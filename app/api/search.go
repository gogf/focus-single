package api

import (
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 搜索管理
var Search = searchApi{}

type searchApi struct{}

// @summary 搜索页面
// @tags    前台-搜索
// @produce html
// @param   key  query string true  "关键字"
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /search [GET]
// @success 200 {string} html "页面HTML"
func (a *searchApi) Index(r *ghttp.Request) {
	var (
		req *internal.ContentSearchReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if searchRes, err := service.Content.Search(r.Context(), req.ContentSearchInput); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			Data:  searchRes,
			Title: service.View.GetTitle(r.Context(), &model.ViewGetTitleInput{}),
		})
	}
}
