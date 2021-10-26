package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/cnt"
	"focus/app/internal/model"
	"focus/app/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// 用户管理
	User = userAct{}
)

type userAct struct{}

func (a *userAct) Index(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	err = a.getContentList(ctx, req)
	return
}

func (a *userAct) Profile(ctx context.Context, req *api.UserProfileReq) (res *api.UserProfileRes, err error) {
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

func (a *userAct) Avatar(ctx context.Context, req *api.UserAvatarReq) (res *api.UserAvatarRes, err error) {
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

func (a *userAct) Password(ctx context.Context, req *api.UserPasswordReq) (res *api.UserPasswordRes, err error) {
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

func (a *userAct) Article(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	req.Type = cnt.ContentTypeArticle
	err = a.getContentList(ctx, req)
	return
}

func (a *userAct) Topic(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	req.Type = cnt.ContentTypeTopic
	err = a.getContentList(ctx, req)
	return
}

func (a *userAct) Ask(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	req.Type = cnt.ContentTypeAsk
	err = a.getContentList(ctx, req)
	return
}

func (a *userAct) getContentList(ctx context.Context, req *api.UserGetContentListReq) error {
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

func (a *userAct) Message(ctx context.Context, req *api.UserGetMessageListReq) (res *api.UserGetMessageListRes, err error) {
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

func (a *userAct) UpdatePassword(ctx context.Context, req *api.UserUpdatePasswordReq) (res *api.UserUpdatePasswordRes, err error) {
	err = service.User.UpdatePassword(ctx, model.UserPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}

func (a *userAct) UpdateAvatar(ctx context.Context, req *api.UserUpdateProfileReq) (res *api.UserUpdateProfileRes, err error) {
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

func (a *userAct) UpdateProfile(ctx context.Context, req *api.UserUpdateProfileReq) (res *api.UserUpdateProfileRes, err error) {
	err = service.User.UpdateProfile(ctx, model.UserUpdateProfileInput{
		Id:       req.Id,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
	return
}

func (a *userAct) Logout(ctx context.Context, req *api.UserLogoutReq) (res *api.UserLogoutRes, err error) {
	if err = service.User.Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(service.Middleware.LoginUrl)
	return
}
