// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package register

import (
	"context"

	"focus-single/api/register/v1"
)

type IRegisterV1 interface {
	RegisterIndex(ctx context.Context, req *v1.RegisterIndexReq) (res *v1.RegisterIndexRes, err error)
	RegisterDo(ctx context.Context, req *v1.RegisterDoReq) (res *v1.RegisterDoRes, err error)
}
