package act

import (
	"context"
	"focus/app/api"
	"focus/app/internal/cnt"
	"focus/app/internal/model"
	"focus/app/internal/service"
	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var (
	// 用户管理
	User = userAct{}
)

type userAct struct{}

// @summary 用户主页
// @tags    前台-用户
// @produce html
// @param   entity body model.UserGetListInput true "请求参数" required
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *userAct) Index(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	err = a.getContentList(ctx, req)
	return
}

// @summary 展示个人资料页面
// @tags    前台-用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
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

// @summary 修改头像页面
// @tags    前台-用户
// @produce html
// @router  /user/avatar [GET]
// @success 200 {string} html "页面HTML"
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

// @summary 修改密码页面
// @tags    前台-用户
// @produce html
// @router  /user/password [GET]
// @success 200 {string} html "页面HTML"
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

// @summary 我的文章页面
// @tags    前台-用户
// @produce html
// @router  /user/article [GET]
// @success 200 {string} html "页面HTML"
func (a *userAct) Article(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	req.Type = cnt.ContentTypeArticle
	err = a.getContentList(ctx, req)
	return
}

// @summary 我的主题页面
// @tags    前台-用户
// @produce html
// @router  /user/topic [GET]
// @success 200 {string} html "页面HTML"
func (a *userAct) Topic(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	req.Type = cnt.ContentTypeTopic
	err = a.getContentList(ctx, req)
	return
}

// @summary 我的问答页面
// @tags    前台-用户
// @produce html
// @router  /user/ask [GET]
// @success 200 {string} html "页面HTML"
func (a *userAct) Ask(ctx context.Context, req *api.UserGetContentListReq) (res *api.UserGetContentListRes, err error) {
	req.Type = cnt.ContentTypeAsk
	err = a.getContentList(ctx, req)
	return
}

// 获取内容列表 参数contentType,用户信息
func (a *userAct) getContentList(ctx context.Context, req *api.UserGetContentListReq) error {
	req.UserId = service.Context.Get(ctx).User.Id
	if output, err := service.User.GetList(ctx, req.UserGetContentListInput); err != nil {
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
func (a *userAct) Message(ctx context.Context, req *api.UserGetMessageListReq) (res *api.UserGetMessageListRes, err error) {
	if getListRes, err := service.User.GetMessageList(ctx, req.UserGetMessageListInput); err != nil {
		return nil, err
	} else {
		service.View.Render(ctx, model.View{
			ContentType: req.TargetType,
			Data:        getListRes,
		})
		return nil, nil
	}
}

// @summary AJAX保存个人资料
// @tags    前台-用户
// @produce json
// @param   entity body internal.UserPasswordReq true "请求参数" required
// @router  /user/update-password [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userAct) UpdatePassword(ctx context.Context, req *api.UserUpdatePasswordReq) (res *api.UserUpdatePasswordRes, err error) {
	err = service.User.UpdatePassword(ctx, req.UserPasswordInput)
	return
}

// @summary AJAX保存个人资料
// @tags    前台-更新头像
// @produce json
// @param   file formData file true "文件域"
// @param   nickname formData string true "请求参数" required
// @router  /user/update-avatar [POST]
// @success 200 {object} response.JsonRes "请求结果"
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
	err = service.User.UpdateAvatar(ctx, req.UserUpdateProfileInput)
	return
}

// @summary AJAX保存个人资料
// @tags    前台-用户
// @produce json
// @param   entity body internal.UserUpdateProfileReq true "请求参数" required
// @router  /user/update-profile [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userAct) UpdateProfile(ctx context.Context, req *api.UserUpdateProfileReq) (res *api.UserUpdateProfileRes, err error) {
	err = service.User.UpdateProfile(ctx, req.UserUpdateProfileInput)
	return
}

// @summary 注销退出
// @description 注销成功后前端引导页面跳转到首页。
// @tags    前台-用户
// @produce json
// @router  /user/logout [GET]
// @success 200 {object} response.JsonRes "执行结果"
func (a *userAct) Logout(ctx context.Context, req *api.UserLogoutReq) (res *api.UserLogoutRes, err error) {
	if err = service.User.Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(service.Middleware.LoginUrl)
	return
}
