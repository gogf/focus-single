package user

import (
	"context"

	"focus-single/api/user/v1"
	"focus-single/internal/consts"
)

func (c *ControllerV1) UserTopic(ctx context.Context, req *v1.UserTopicReq) (res *v1.UserTopicRes, err error) {
	req.Type = consts.ContentTypeTopic
	err = c.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}
