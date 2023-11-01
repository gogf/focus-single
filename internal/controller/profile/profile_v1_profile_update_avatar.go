package profile

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"focus-single/api/profile/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ProfileUpdateAvatar(ctx context.Context, req *v1.ProfileUpdateAvatarReq) (res *v1.ProfileUpdateAvatarRes, err error) {
	var (
		request    = g.RequestFromCtx(ctx)
		uploadFile = request.GetUploadFile("file")
	)
	if uploadFile == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	uploadResult, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       uploadFile,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	if uploadResult != nil {
		req.Avatar = uploadResult.Url
	}
	err = service.User().UpdateAvatar(ctx, model.UserUpdateAvatarInput{
		UserId: service.BizCtx().Get(ctx).User.Id,
		Avatar: req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	// 更新登录session Avatar
	sessionProfile := service.Session().GetUser(ctx)
	sessionProfile.Avatar = req.Avatar
	err = service.Session().SetUser(ctx, sessionProfile)
	return
}
