package profile

import (
	"context"

	"focus-single/api/profile/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ProfileUpdatePassword(ctx context.Context, req *v1.ProfileUpdatePasswordReq) (res *v1.ProfileUpdatePasswordRes, err error) {
	err = service.User().UpdatePassword(ctx, model.UserPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}
