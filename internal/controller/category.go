package controller

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/service"
)

// 栏目管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Tree(ctx context.Context, req *apiv1.CategoryTreeReq) (res *apiv1.CategoryTreeRes, err error) {
	res = &apiv1.CategoryTreeRes{}
	res.List, err = service.Category().GetTree(ctx, req.ContentType)
	return
}
