package user

import (
	"context"

	"focus-single/api/user/v1"
	"focus-single/internal/consts"
)

func (c *ControllerV1) UserArticle(ctx context.Context, req *v1.UserArticleReq) (res *v1.UserArticleRes, err error) {
	req.Type = consts.ContentTypeArticle
	err = c.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}
