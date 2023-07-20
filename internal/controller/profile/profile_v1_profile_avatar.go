package profile

import (
	"context"

	"focus-single/api/profile/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ProfileAvatar(ctx context.Context, req *v1.ProfileAvatarReq) (res *v1.ProfileAvatarRes, err error) {
	out, err := service.User().GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 头像"
	service.View().Render(ctx, model.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}
