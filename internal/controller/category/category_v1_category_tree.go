package category

import (
	"context"

	"focus-single/api/category/v1"
	"focus-single/internal/service"
)

func (c *ControllerV1) CategoryTree(ctx context.Context, req *v1.CategoryTreeReq) (res *v1.CategoryTreeRes, err error) {
	res = &v1.CategoryTreeRes{}
	res.List, err = service.Category().GetTree(ctx, req.ContentType)
	return
}
