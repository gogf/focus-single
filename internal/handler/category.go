package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/service"
)

var (
	// 栏目管理
	Category = hCategory{}
)

type hCategory struct{}

func (a *hCategory) Tree(ctx context.Context, req *apiv1.CategoryTreeReq) (res *apiv1.CategoryTreeRes, err error) {
	res = &apiv1.CategoryTreeRes{}
	res.List, err = service.Category().GetTree(ctx, req.ContentType)
	return
}
