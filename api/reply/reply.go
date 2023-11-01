// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package reply

import (
	"context"

	"focus-single/api/reply/v1"
)

type IReplyV1 interface {
	ReplyGetListContent(ctx context.Context, req *v1.ReplyGetListContentReq) (res *v1.ReplyGetListContentRes, err error)
	ReplyCreate(ctx context.Context, req *v1.ReplyCreateReq) (res *v1.ReplyCreateRes, err error)
	ReplyDelete(ctx context.Context, req *v1.ReplyDeleteReq) (res *v1.ReplyDeleteRes, err error)
}
