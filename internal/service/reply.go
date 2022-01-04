package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"

	"focus-single/internal/model"
	"focus-single/internal/model/entity"
	"focus-single/internal/service/internal/dao"
)

type sReply struct{}

var insReply = sReply{}

// Reply 评论/回复管理服务
func Reply() *sReply {
	return &insReply
}

// Create 创建回复
func (s *sReply) Create(ctx context.Context, in model.ReplyCreateInput) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 覆盖用户ID
		in.UserId = Context().Get(ctx).User.Id
		_, err := dao.Reply.Ctx(ctx).Data(in).Insert()
		if err == nil {
			err = Content().AddReplyCount(ctx, in.TargetId, 1)
		}
		return err
	})
}

// Delete 删除回复(硬删除)
func (s *sReply) Delete(ctx context.Context, id uint) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var reply *entity.Reply
		err := dao.Reply.Ctx(ctx).WherePri(id).Scan(&reply)
		if err != nil {
			return err
		}
		// 删除回复记录
		_, err = dao.Reply.Ctx(ctx).Where(g.Map{
			dao.Reply.Columns().Id:     id,
			dao.Reply.Columns().UserId: Context().Get(ctx).User.Id,
		}).Delete()
		if err == nil {
			// 回复统计-1
			err = Content().AddReplyCount(ctx, reply.TargetId, -1)
			if err != nil {
				return err
			}
			// 判断回复是否采纳
			var content *entity.Content
			err = dao.Content.Ctx(ctx).WherePri(reply.TargetId).Scan(&content)
			if err == nil && content != nil && content.AdoptedReplyId == id {
				err = Content().UnacceptedReply(ctx, reply.TargetId)
			}
		}
		return err
	})
}

// 删除回复(硬删除)
func (s *sReply) DeleteByUserContentId(ctx context.Context, userId, contentId uint) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除内容对应的回复
		_, err := dao.Reply.Ctx(ctx).Where(g.Map{
			dao.Reply.Columns().TargetId: contentId,
			dao.Reply.Columns().UserId:   userId,
		}).Delete()
		return err
	})
}

// 获取回复列表
func (s *sReply) GetList(ctx context.Context, in model.ReplyGetListInput) (out *model.ReplyGetListOutput, err error) {
	out = &model.ReplyGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	m := dao.Reply.Ctx(ctx).Fields(model.ReplyListItem{})
	if in.TargetType != "" {
		m = m.Where(dao.Reply.Columns().TargetType, in.TargetType)
	}
	if in.TargetId > 0 {
		m = m.Where(dao.Reply.Columns().TargetId, in.TargetId)
	}
	if in.UserId > 0 {
		m = m.Where(dao.Reply.Columns().UserId, in.UserId)
	}

	err = m.Page(in.Page, in.Size).OrderDesc(dao.Content.Columns().Id).ScanList(&out.List, "Reply")
	if err != nil {
		return nil, err
	}
	if len(out.List) == 0 {
		return nil, nil
	}
	// User
	if err = m.ScanList(&out.List, "Reply"); err != nil {
		return nil, err
	}
	err = dao.User.Ctx(ctx).
		Fields(model.ReplyListUserItem{}).
		Where(dao.User.Columns().Id, gutil.ListItemValuesUnique(out.List, "Reply", "UserId")).
		ScanList(&out.List, "User", "Reply", "id:UserId")
	if err != nil {
		return nil, err
	}

	// Content
	err = dao.Content.Ctx(ctx).
		Fields(dao.Content.Columns().Id, dao.Content.Columns().Title, dao.Content.Columns().CategoryId).
		Where(dao.Content.Columns().Id, gutil.ListItemValuesUnique(out.List, "Reply", "TargetId")).
		ScanList(&out.List, "Content", "Reply", "id:TargetId")
	if err != nil {
		return nil, err
	}

	// Category
	err = dao.Category.Ctx(ctx).
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns().Id, gutil.ListItemValuesUnique(out.List, "Content", "CategoryId")).
		ScanList(&out.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}

	return out, nil
}
