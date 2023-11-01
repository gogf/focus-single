// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package interact

import (
	"context"

	"focus-single/api/interact/v1"
)

type IInteractV1 interface {
	InteractZan(ctx context.Context, req *v1.InteractZanReq) (res *v1.InteractZanRes, err error)
	InteractCancelZan(ctx context.Context, req *v1.InteractCancelZanReq) (res *v1.InteractCancelZanRes, err error)
	InteractCai(ctx context.Context, req *v1.InteractCaiReq) (res *v1.InteractCaiRes, err error)
	InteractCancelCai(ctx context.Context, req *v1.InteractCancelCaiReq) (res *v1.InteractCancelCaiRes, err error)
}
