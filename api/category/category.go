// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package category

import (
	"context"

	"focus-single/api/category/v1"
)

type ICategoryV1 interface {
	CategoryTree(ctx context.Context, req *v1.CategoryTreeReq) (res *v1.CategoryTreeRes, err error)
}
