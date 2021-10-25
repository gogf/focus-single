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

func (a *categoryAct) Tree(ctx context.Context, req *api.CategoryGetTreeReq) (res *api.CategoryGetTreeRes, err error) {
	res = &api.CategoryGetTreeRes{}
	res.List, err = service.Category.GetTree(ctx, req.ContentType)
	return
}
