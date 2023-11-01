package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"focus-single/api/user/v1"
	"focus-single/internal/consts"
	"focus-single/internal/service"
)

func (c *ControllerV1) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	if err = service.User().Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(consts.UserLoginUrl)
	return
}
