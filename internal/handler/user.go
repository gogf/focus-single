package handler

import (
	"context"

	"focus-single/apiv1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// 用户管理
	User = handlerUser{}
)

type handlerUser struct{}

func (a *handlerUser) Index(ctx context.Context, req *apiv1.UserGetContentListReq) (res *apiv1.UserGetContentListRes, err error) {
	err = a.getContentList(ctx, req)
	return
}

func (a *handlerUser) Profile(ctx context.Context, req *apiv1.UserProfileReq) (res *apiv1.UserProfileRes, err error) {
	if getProfile, err := service.User.GetProfile(ctx); err != nil {
		return nil, err
	} else {
		title := "用户 " + getProfile.Nickname + " 资料"
		service.View.Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil, nil
	}
}

func (a *handlerUser) Avatar(ctx context.Context, req *apiv1.UserAvatarReq) (res *apiv1.UserAvatarRes, err error) {
	if getProfile, err := service.User.GetProfile(ctx); err != nil {
		return nil, err
	} else {
		title := "用户 " + getProfile.Nickname + " 头像"
		service.View.Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil, nil
	}
}

func (a *handlerUser) Password(ctx context.Context, req *apiv1.UserPasswordReq) (res *apiv1.UserPasswordRes, err error) {
	if getProfile, err := service.User.GetProfile(ctx); err != nil {
		return nil, err
	} else {
		title := "用户 " + getProfile.Nickname + " 修改密码"
		service.View.Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil, nil
	}
}

func (a *handlerUser) Article(ctx context.Context, req *apiv1.UserGetContentListReq) (res *apiv1.UserGetContentListRes, err error) {
	req.Type = consts.ContentTypeArticle
	err = a.getContentList(ctx, req)
	return
}

func (a *handlerUser) Topic(ctx context.Context, req *apiv1.UserGetContentListReq) (res *apiv1.UserGetContentListRes, err error) {
	req.Type = consts.ContentTypeTopic
	err = a.getContentList(ctx, req)
	return
}

func (a *handlerUser) Ask(ctx context.Context, req *apiv1.UserGetContentListReq) (res *apiv1.UserGetContentListRes, err error) {
	req.Type = consts.ContentTypeAsk
	err = a.getContentList(ctx, req)
	return
}

func (a *handlerUser) getContentList(ctx context.Context, req *apiv1.UserGetContentListReq) error {
	req.UserId = service.Context.Get(ctx).User.Id
	if out, err := service.User.GetList(ctx, model.UserGetContentListInput{
		ContentGetListInput: model.ContentGetListInput{
			Type:       req.Type,
			CategoryId: req.CategoryId,
			Page:       req.Page,
			Size:       req.Size,
			Sort:       req.Sort,
			UserId:     req.UserId,
		},
	}); err != nil {
		return err
	} else {
		title := service.View.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		})
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        out,
			Title:       title,
		})
		return nil
	}
}

func (a *handlerUser) Message(ctx context.Context, req *apiv1.UserGetMessageListReq) (res *apiv1.UserGetMessageListRes, err error) {
	if getListRes, err := service.User.GetMessageList(ctx, model.UserGetMessageListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		UserId:     service.Session.GetUser(ctx).Id,
	}); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: req.TargetType,
			Data:        getListRes,
		})
		return nil, nil
	}
}

func (a *handlerUser) UpdatePassword(ctx context.Context, req *apiv1.UserUpdatePasswordReq) (res *apiv1.UserUpdatePasswordRes, err error) {
	err = service.User.UpdatePassword(ctx, model.UserPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}

func (a *handlerUser) UpdateAvatar(ctx context.Context, req *apiv1.UserUpdateProfileReq) (res *apiv1.UserUpdateProfileRes, err error) {
	var (
		request = g.RequestFromCtx(ctx)
		file    = request.GetUploadFile("file")
	)
	if file == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	uploadResult, err := service.File.Upload(ctx, model.FileUploadInput{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	if uploadResult != nil {
		req.Avatar = uploadResult.Url
	}
	err = service.User.UpdateAvatar(ctx, model.UserUpdateProfileInput{
		Id:       req.Id,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
	return
}

func (a *handlerUser) UpdateProfile(ctx context.Context, req *apiv1.UserUpdateProfileReq) (res *apiv1.UserUpdateProfileRes, err error) {
	err = service.User.UpdateProfile(ctx, model.UserUpdateProfileInput{
		Id:       req.Id,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
	return
}

func (a *handlerUser) Logout(ctx context.Context, req *apiv1.UserLogoutReq) (res *apiv1.UserLogoutRes, err error) {
	if err = service.User.Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(service.Middleware.LoginUrl)
	return
}
