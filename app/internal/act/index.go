package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/model"
	"focus/app/internal/service"
)

var (
	// 首页接口
	Index = indexAct{}
)

type indexAct struct{}

// @summary 展示站点首页
// @tags    前台-首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func (a *indexAct) Index(ctx context.Context, req *api.ContentGetListReq) (res *api.ContentGetListRes, err error) {
	if getListRes, err := service.Content.GetList(ctx, req.ContentGetListInput); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       "首页",
		})
	}
	return
}
