// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package login

import (
	"context"

	"focus-single/api/login/v1"
)

type ILoginV1 interface {
	LoginIndex(ctx context.Context, req *v1.LoginIndexReq) (res *v1.LoginIndexRes, err error)
	LoginDo(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error)
}
