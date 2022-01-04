package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/service"
)

var (
	// Category 栏目管理
	Category = handlerCategory{}
)

type handlerCategory struct{}

func (a *handlerCategory) Tree(ctx context.Context, req *apiv1.CategoryTreeReq) (res *apiv1.CategoryTreeRes, err error) {
	res = &apiv1.CategoryTreeRes{}
	res.List, err = service.Category().GetTree(ctx, req.ContentType)
	return
}
