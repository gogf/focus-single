package user

import (
	"context"

	"focus-single/api/user/v1"
	"focus-single/internal/consts"
)

func (c *ControllerV1) UserAsk(ctx context.Context, req *v1.UserAskReq) (res *v1.UserAskRes, err error) {
	req.Type = consts.ContentTypeAsk
	err = c.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}
