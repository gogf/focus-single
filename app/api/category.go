package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/model"
	"focus/app/service"
)

// 栏目管理
var Category = categoryApi{}

type categoryApi struct{}

// @summary 获取分类列表，构造成树形结构返回
// @tags    前台-分类
// @produce json
// @param   contentType query string true  "分类类型:topic, ask, article, reply，当传递空时表示获取所有栏目"
// @router  /category/tree [GET]
// @success 200 {array} model.CategoryTreeItem "分类列表"
func (a *categoryApi) Tree(ctx context.Context, req *internal.CategoryGetTreeReq) (res []*model.CategoryTreeItem, err error) {
	return service.Category.GetTree(ctx, req.ContentType)
}
