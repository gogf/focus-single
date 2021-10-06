package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/service"
)

var (
	// 栏目管理
	Category = categoryAct{}
)

type categoryAct struct{}

// @summary 获取分类列表，构造成树形结构返回
// @tags    前台-分类
// @produce json
// @param   contentType query string true  "分类类型:topic, ask, article, reply，当传递空时表示获取所有栏目"
// @router  /category/tree [GET]
// @success 200 {array} model.CategoryTreeItem "分类列表"
func (a *categoryAct) Tree(ctx context.Context, req *api.CategoryGetTreeReq) (res *api.CategoryGetTreeRes, err error) {
	res = &api.CategoryGetTreeRes{}
	res.List, err = service.Category.GetTree(ctx, req.ContentType)
	return
}
