package category

import (
	"context"

	v1 "focus-single/api/v1/category"
	"focus-single/internal/service/category"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Tree(ctx context.Context, req *v1.TreeReq) (res *v1.TreeRes, err error) {
	res = &v1.TreeRes{}
	res.List, err = category.GetTree(ctx, req.ContentType)
	return
}
