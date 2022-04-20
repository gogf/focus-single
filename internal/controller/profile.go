package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/service/bizctx"
	"focus-single/internal/service/file"
	"focus-single/internal/service/reply"
	"focus-single/internal/service/session"
	"focus-single/internal/service/user"
	"focus-single/internal/service/view"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 个人中心
var Profile = cProfile{}

type cProfile struct{}

func (a *cProfile) Index(ctx context.Context, req *v1.ProfileIndexReq) (res *v1.ProfileIndexRes, err error) {
	out, err := user.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 资料"
	view.Render(ctx, view.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (a *cProfile) Avatar(ctx context.Context, req *v1.ProfileAvatarReq) (res *v1.ProfileAvatarRes, err error) {
	out, err := user.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 头像"
	view.Render(ctx, view.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (a *cProfile) UpdateAvatar(ctx context.Context, req *v1.ProfileUpdateAvatarReq) (res *v1.ProfileUpdateAvatarRes, err error) {
	var (
		request    = g.RequestFromCtx(ctx)
		uploadFile = request.GetUploadFile("file")
	)
	if uploadFile == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	uploadResult, err := file.Upload(ctx, file.UploadInput{
		File:       uploadFile,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	if uploadResult != nil {
		req.Avatar = uploadResult.Url
	}
	err = user.UpdateAvatar(ctx, user.UpdateAvatarInput{
		UserId: bizctx.Get(ctx).User.Id,
		Avatar: req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	// 更新登录session Avatar
	sessionProfile := session.GetUser(ctx)
	sessionProfile.Avatar = req.Avatar
	err = session.SetUser(ctx, sessionProfile)
	return
}

func (a *cProfile) Password(ctx context.Context, req *v1.ProfilePasswordReq) (res *v1.ProfilePasswordRes, err error) {
	out, err := user.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 修改密码"
	view.Render(ctx, view.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (a *cProfile) UpdatePassword(ctx context.Context, req *v1.ProfileUpdatePasswordReq) (res *v1.ProfileUpdatePasswordRes, err error) {
	err = user.UpdatePassword(ctx, user.PasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}

func (a *cProfile) UpdateProfile(ctx context.Context, req *v1.ProfileUpdateReq) (res *v1.ProfileUpdateRes, err error) {
	err = user.UpdateProfile(ctx, user.UpdateProfileInput{
		UserId:   req.Id,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
	return
}

func (a *cProfile) Message(ctx context.Context, req *v1.ProfileMessageReq) (res *v1.ProfileMessageRes, err error) {
	type ViewData struct {
		List  []reply.GetListOutputItem // 列表
		Page  int                       // 分页码
		Size  int                       // 分页数量
		Total int                       // 数据总数
		Stats map[string]int            // 发布内容数量
	}
	var (
		ctxUser = bizctx.Get(ctx).User
		in      = reply.GetListInput{
			Page:       req.Page,
			Size:       req.Size,
			TargetType: req.TargetType,
			TargetId:   req.TargetId,
		}
	)
	if !ctxUser.IsAdmin {
		in.UserId = ctxUser.Id
	}
	// 回复列表
	replyListOut, err := reply.GetList(ctx, in)
	if err != nil {
		return nil, err
	}
	var data = ViewData{
		Page:  req.Page,
		Size:  req.Size,
		List:  replyListOut.List,
		Total: replyListOut.Total,
	}
	if err != nil {
		return nil, err
	}
	// 用户信息统计
	data.Stats, err = user.GetUserStats(ctx, ctxUser.Id)
	if err != nil {
		return nil, err
	}
	view.Render(ctx, view.View{
		ContentType: req.TargetType,
		Data:        data,
	})
	return nil, nil
}
