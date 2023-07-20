// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package search

import (
	"context"

	"focus-single/api/search/v1"
)

type ISearchV1 interface {
	SearchIndex(ctx context.Context, req *v1.SearchIndexReq) (res *v1.SearchIndexRes, err error)
}
