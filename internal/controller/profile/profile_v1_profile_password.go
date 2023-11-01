package profile

import (
	"context"

	"focus-single/api/profile/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ProfilePassword(ctx context.Context, req *v1.ProfilePasswordReq) (res *v1.ProfilePasswordRes, err error) {
	out, err := service.User().GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 修改密码"
	service.View().Render(ctx, model.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}
