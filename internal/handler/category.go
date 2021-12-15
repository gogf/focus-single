package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/service"
)

var (
	// 栏目管理
	Category = handlerCategory{}
)

type handlerCategory struct{}

func (a *handlerCategory) Tree(ctx context.Context, req *apiv1.CategoryGetTreeReq) (res *apiv1.CategoryGetTreeRes, err error) {
	res = &apiv1.CategoryGetTreeRes{}
	res.List, err = service.Category.GetTree(ctx, req.ContentType)
	return
}
