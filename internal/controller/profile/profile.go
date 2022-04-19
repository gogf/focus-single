package profile

import (
	"context"

	v1 "focus-single/api/v1/profile"
	"focus-single/internal/model"
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

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	out, err := user.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 资料"
	view.Render(ctx, model.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (c *controller) Avatar(ctx context.Context, req *v1.AvatarReq) (res *v1.AvatarRes, err error) {
	out, err := user.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 头像"
	view.Render(ctx, model.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (c *controller) UpdateAvatar(ctx context.Context, req *v1.UpdateAvatarReq) (res *v1.UpdateAvatarRes, err error) {
	var (
		request    = g.RequestFromCtx(ctx)
		uploadFile = request.GetUploadFile("file")
	)
	if uploadFile == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	uploadResult, err := file.Upload(ctx, model.FileUploadInput{
		File:       uploadFile,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	if uploadResult != nil {
		req.Avatar = uploadResult.Url
	}
	err = user.UpdateAvatar(ctx, model.UserUpdateAvatarInput{
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

func (c *controller) Password(ctx context.Context, req *v1.PasswordReq) (res *v1.PasswordRes, err error) {
	out, err := user.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 修改密码"
	view.Render(ctx, model.View{
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (c *controller) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error) {
	err = user.UpdatePassword(ctx, model.UserPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}

func (c *controller) UpdateProfile(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = user.UpdateProfile(ctx, model.UserUpdateProfileInput{
		UserId:   req.Id,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
	return
}

func (c *controller) Message(ctx context.Context, req *v1.MessageReq) (res *v1.MessageRes, err error) {
	type ViewData struct {
		List  []model.ReplyGetListOutputItem // 列表
		Page  int                            // 分页码
		Size  int                            // 分页数量
		Total int                            // 数据总数
		Stats map[string]int                 // 发布内容数量
	}
	var (
		ctxUser = bizctx.Get(ctx).User
		in      = model.ReplyGetListInput{
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
	view.Render(ctx, model.View{
		ContentType: req.TargetType,
		Data:        data,
	})
	return nil, nil
}