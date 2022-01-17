package controller

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/model"
	"focus-single/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 个人中心
var Profile = cProfile{}

type cProfile struct{}

func (a *cProfile) Index(ctx context.Context, req *apiv1.ProfileIndexReq) (res *apiv1.ProfileIndexRes, err error) {
	if getProfile, err := service.User().GetProfile(ctx); err != nil {
		return nil, err
	} else {
		title := "用户 " + getProfile.Nickname + " 资料"
		service.View().Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil, nil
	}
}

func (a *cProfile) Avatar(ctx context.Context, req *apiv1.ProfileAvatarReq) (res *apiv1.ProfileAvatarRes, err error) {
	if getProfile, err := service.User().GetProfile(ctx); err != nil {
		return nil, err
	} else {
		title := "用户 " + getProfile.Nickname + " 头像"
		service.View().Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil, nil
	}
}

func (a *cProfile) UpdateAvatar(ctx context.Context, req *apiv1.ProfileUpdateAvatarReq) (res *apiv1.ProfileUpdateAvatarRes, err error) {
	var (
		request = g.RequestFromCtx(ctx)
		file    = request.GetUploadFile("file")
	)
	if file == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	uploadResult, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	if uploadResult != nil {
		req.Avatar = uploadResult.Url
	}
	err = service.User().UpdateAvatar(ctx, model.UserUpdateAvatarInput{
		UserId: service.Context().Get(ctx).User.Id,
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

func (a *cProfile) Password(ctx context.Context, req *apiv1.ProfilePasswordReq) (res *apiv1.ProfilePasswordRes, err error) {
	if getProfile, err := service.User().GetProfile(ctx); err != nil {
		return nil, err
	} else {
		title := "用户 " + getProfile.Nickname + " 修改密码"
		service.View().Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil, nil
	}
}

func (a *cProfile) UpdatePassword(ctx context.Context, req *apiv1.ProfileUpdatePasswordReq) (res *apiv1.ProfileUpdatePasswordRes, err error) {
	err = service.User().UpdatePassword(ctx, model.UserPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}

func (a *cProfile) UpdateProfile(ctx context.Context, req *apiv1.ProfileUpdateReq) (res *apiv1.ProfileUpdateRes, err error) {
	err = service.User().UpdateProfile(ctx, model.UserUpdateProfileInput{
		UserId:   req.Id,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
	return
}

func (a *cProfile) Message(ctx context.Context, req *apiv1.ProfileMessageReq) (res *apiv1.ProfileMessageRes, err error) {
	if getListRes, err := service.User().GetMessageList(ctx, model.UserGetMessageListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		UserId:     service.Session().GetUser(ctx).Id,
	}); err != nil {
		return nil, err
	} else {
		service.View().Render(ctx, model.View{
			ContentType: req.TargetType,
			Data:        getListRes,
		})
		return nil, nil
	}
}
