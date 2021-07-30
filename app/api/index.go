package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
)

// 首页接口
var Index = indexApi{}

type indexApi struct{}

// @summary 展示站点首页
// @tags    前台-首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func (a *indexApi) Index(ctx context.Context, req *internal.ContentGetListReq) error {
	if getListRes, err := service.Content.GetList(ctx, req.ContentGetListInput); err != nil {
		return err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       "首页",
		})
	}
	return nil
}
