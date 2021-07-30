package api

import (
	"context"
	"focus/app/api/internal"
	"focus/app/cnt"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var User = userApi{}

type userApi struct{}

// @summary 用户主页
// @tags    前台-用户
// @produce html
// @param   entity body model.UserGetListInput true "请求参数" required
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(ctx context.Context, req *internal.UserGetListReq) error {
	return a.getContentList(ctx, req)
}

// @summary 展示个人资料页面
// @tags    前台-用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Profile(ctx context.Context) error {
	if getProfile, err := service.User.GetProfile(ctx); err != nil {
		return err
	} else {
		title := "用户 " + getProfile.Nickname + " 资料"
		service.View.Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil
	}
}

// @summary 修改头像页面
// @tags    前台-用户
// @produce html
// @router  /user/avatar [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Avatar(ctx context.Context) error {
	if getProfile, err := service.User.GetProfile(ctx); err != nil {
		return err
	} else {
		title := "用户 " + getProfile.Nickname + " 头像"
		service.View.Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil
	}
}

// @summary 修改密码页面
// @tags    前台-用户
// @produce html
// @router  /user/password [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Password(ctx context.Context) error {
	if getProfile, err := service.User.GetProfile(ctx); err != nil {
		return err
	} else {
		title := "用户 " + getProfile.Nickname + " 修改密码"
		service.View.Render(ctx, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
		return nil
	}
}

// @summary 我的文章页面
// @tags    前台-用户
// @produce html
// @router  /user/article [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Article(ctx context.Context, req *internal.UserGetListReq) error {
	req.Type = cnt.ContentTypeArticle
	return a.getContentList(ctx, req)
}

// @summary 我的主题页面
// @tags    前台-用户
// @produce html
// @router  /user/topic [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Topic(ctx context.Context, req *internal.UserGetListReq) error {
	req.Type = cnt.ContentTypeTopic
	return a.getContentList(ctx, req)
}

// @summary 我的问答页面
// @tags    前台-用户
// @produce html
// @router  /user/ask [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Ask(ctx context.Context, req *internal.UserGetListReq) error {
	req.Type = cnt.ContentTypeAsk
	return a.getContentList(ctx, req)
}

// 获取内容列表 参数contentType,用户信息
func (a *userApi) getContentList(ctx context.Context, req *internal.UserGetListReq) error {
	req.UserId = service.Context.Get(ctx).User.Id
	if output, err := service.User.GetList(ctx, req.UserGetListInput); err != nil {
		return err
	} else {
		title := service.View.GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		})
		service.View.Render(ctx, model.View{
			ContentType: req.Type,
			Data:        output,
			Title:       title,
		})
		return nil
	}
}

// @summary 我的消息页面
// @tags    前台-用户
// @produce html
// @router  /user/message [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Message(ctx context.Context, req *internal.UserGetMessageListReq) error {
	if getListRes, err := service.User.GetMessageList(ctx, req.UserGetMessageListInput); err != nil {
		return err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: req.TargetType,
			Data:        getListRes,
		})
		return nil
	}
}

// @summary AJAX保存个人资料
// @tags    前台-用户
// @produce json
// @param   entity body internal.UserPasswordReq true "请求参数" required
// @router  /user/update-password [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdatePassword(ctx context.Context, req *internal.UserPasswordReq) error {
	return service.User.UpdatePassword(ctx, req.UserPasswordInput)
}

// @summary AJAX保存个人资料
// @tags    前台-更新头像
// @produce json
// @param   file formData file true "文件域"
// @param   nickname formData string true "请求参数" required
// @router  /user/update-avatar [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdateAvatar(ctx context.Context, req *internal.UserUpdateProfileReq) error {
	var (
		request = g.RequestFromCtx(ctx)
		file    = request.GetUploadFile("file")
	)
	if file == nil {
		return gerror.NewCode(gerror.CodeMissingParameter, "请选择需要上传的文件")
	}
	uploadResult, err := service.File.Upload(ctx, model.FileUploadInput{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		return err
	}
	if uploadResult != nil {
		req.Avatar = uploadResult.Url
	}
	return service.User.UpdateAvatar(ctx, req.UserUpdateProfileInput)
}

// @summary AJAX保存个人资料
// @tags    前台-用户
// @produce json
// @param   entity body internal.UserUpdateProfileReq true "请求参数" required
// @router  /user/update-profile [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdateProfile(ctx context.Context, req *internal.UserUpdateProfileReq) error {
	return service.User.UpdateProfile(ctx, req.UserUpdateProfileInput)
}

// @summary 注销退出
// @description 注销成功后前端引导页面跳转到首页。
// @tags    前台-用户
// @produce json
// @router  /user/logout [GET]
// @success 200 {object} response.JsonRes "执行结果"
func (a *userApi) Logout(ctx context.Context) error {
	if err := service.User.Logout(ctx); err != nil {
		return err
	} else {
		g.RequestFromCtx(ctx).Response.RedirectTo(service.Middleware.LoginUrl)
		return nil
	}
}
