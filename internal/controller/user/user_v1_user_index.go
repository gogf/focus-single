package user

import (
	"context"

	"focus-single/api/user/v1"
)

func (c *ControllerV1) UserIndex(ctx context.Context, req *v1.UserIndexReq) (res *v1.UserIndexRes, err error) {
	err = c.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}
