// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package content

import (
	"context"

	"focus-single/api/content/v1"
)

type IContentV1 interface {
	ContentGetListCommon(ctx context.Context, req *v1.ContentGetListCommonReq) (res *v1.ContentGetListCommonRes, err error)
	ContentShowCreate(ctx context.Context, req *v1.ContentShowCreateReq) (res *v1.ContentShowCreateRes, err error)
	ContentCreate(ctx context.Context, req *v1.ContentCreateReq) (res *v1.ContentCreateRes, err error)
	ContentShowUpdate(ctx context.Context, req *v1.ContentShowUpdateReq) (res *v1.ContentShowUpdateRes, err error)
	ContentUpdate(ctx context.Context, req *v1.ContentUpdateReq) (res *v1.ContentUpdateRes, err error)
	ContentDelete(ctx context.Context, req *v1.ContentDeleteReq) (res *v1.ContentDeleteRes, err error)
}
