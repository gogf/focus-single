// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"focus-single/api/user/v1"
)

type IUserV1 interface {
	UserIndex(ctx context.Context, req *v1.UserIndexReq) (res *v1.UserIndexRes, err error)
	UserArticle(ctx context.Context, req *v1.UserArticleReq) (res *v1.UserArticleRes, err error)
	UserTopic(ctx context.Context, req *v1.UserTopicReq) (res *v1.UserTopicRes, err error)
	UserAsk(ctx context.Context, req *v1.UserAskReq) (res *v1.UserAskRes, err error)
	UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error)
}
