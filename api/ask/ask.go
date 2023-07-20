// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ask

import (
	"context"

	"focus-single/api/ask/v1"
)

type IAskV1 interface {
	AskIndex(ctx context.Context, req *v1.AskIndexReq) (res *v1.AskIndexRes, err error)
	AskDetail(ctx context.Context, req *v1.AskDetailReq) (res *v1.AskDetailRes, err error)
}
