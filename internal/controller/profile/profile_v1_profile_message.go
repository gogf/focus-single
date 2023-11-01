package profile

import (
	"context"

	"focus-single/api/profile/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) ProfileMessage(ctx context.Context, req *v1.ProfileMessageReq) (res *v1.ProfileMessageRes, err error) {
	type ViewData struct {
		List  []model.ReplyGetListOutputItem // 列表
		Page  int                            // 分页码
		Size  int                            // 分页数量
		Total int                            // 数据总数
		Stats map[string]int                 // 发布内容数量
	}
	var (
		ctxUser = service.BizCtx().Get(ctx).User
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
	replyListOut, err := service.Reply().GetList(ctx, in)
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
	data.Stats, err = service.User().GetUserStats(ctx, ctxUser.Id)
	if err != nil {
		return nil, err
	}
	service.View().Render(ctx, model.View{
		ContentType: req.TargetType,
		Data:        data,
	})
	return nil, nil
}
