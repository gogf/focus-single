package login

import (
	"context"

	"focus-single/api/login/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) LoginIndex(ctx context.Context, req *v1.LoginIndexReq) (res *v1.LoginIndexRes, err error) {
	service.View().Render(ctx, model.View{})
	return
}
