// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"focus-single/api/article/v1"
)

type IArticleV1 interface {
	ArticleIndex(ctx context.Context, req *v1.ArticleIndexReq) (res *v1.ArticleIndexRes, err error)
	ArticleDetail(ctx context.Context, req *v1.ArticleDetailReq) (res *v1.ArticleDetailRes, err error)
}
