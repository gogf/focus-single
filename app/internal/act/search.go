package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 搜索管理
	Search = searchAct{}
)

type searchAct struct{}

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
func (a *searchAct) Index(ctx context.Context, req *api.ContentSearchReq) (res *api.ContentSearchRes, err error) {
	if searchRes, err := service.Content.Search(ctx, req.ContentSearchInput); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			Data:  searchRes,
			Title: service.View.GetTitle(ctx, &model.ViewGetTitleInput{}),
		})
		return nil, nil
	}
}
