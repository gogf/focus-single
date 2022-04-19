package reply

import (
	"context"

	"focus-single/internal/model/do"
	"focus-single/internal/service/bizctx"
	"focus-single/internal/service/content"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"

	"focus-single/internal/dao"
	"focus-single/internal/model"
	"focus-single/internal/model/entity"
)

// Create 创建回复
func Create(ctx context.Context, in model.ReplyCreateInput) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 覆盖用户ID
		in.UserId = bizctx.Get(ctx).User.Id
		_, err := dao.Reply.Ctx(ctx).Data(in).Insert()
		if err == nil {
			err = content.AddReplyCount(ctx, in.TargetId, 1)
		}
		return err
	})
}

// Delete 删除回复(硬删除)
func Delete(ctx context.Context, id uint) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var reply *entity.Reply
		err := dao.Reply.Ctx(ctx).WherePri(id).Scan(&reply)
		if err != nil {
			return err
		}
		// 删除回复记录
		_, err = dao.Reply.Ctx(ctx).Where(do.Reply{
			Id:     id,
			UserId: bizctx.Get(ctx).User.Id,
		}).Delete()
		if err == nil {
			// 回复统计-1
			err = content.AddReplyCount(ctx, reply.TargetId, -1)
			if err != nil {
				return err
			}
			// 判断回复是否采纳
			var contentEntity *entity.Content
			err = dao.Content.Ctx(ctx).WherePri(reply.TargetId).Scan(&contentEntity)
			if err == nil && contentEntity != nil && contentEntity.AdoptedReplyId == id {
				err = content.UnacceptedReply(ctx, reply.TargetId)
			}
		}
		return err
	})
}

// 获取回复列表
func GetList(ctx context.Context, in model.ReplyGetListInput) (out *model.ReplyGetListOutput, err error) {
	out = &model.ReplyGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	var (
		orm = dao.Reply.Ctx(ctx)
		cls = dao.Reply.Columns()
	)
	if in.TargetType != "" {
		orm = orm.Where(cls.TargetType, in.TargetType)
	}
	if in.TargetId > 0 {
		orm = orm.Where(cls.TargetId, in.TargetId)
	}
	if in.UserId > 0 {
		orm = orm.Where(cls.UserId, in.UserId)
	}
	err = orm.Page(in.Page, in.Size).OrderDesc(cls.Id).ScanList(&out.List, "Reply")
	if err != nil {
		return nil, err
	}
	if len(out.List) == 0 {
		return nil, nil
	}
	// User
	if err = orm.ScanList(&out.List, "Reply"); err != nil {
		return nil, err
	}
	var userIds = gutil.ListItemValuesUnique(out.List, "Reply", "UserId")
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userIds).
		ScanList(&out.List, "User", "Reply", "id:UserId")
	if err != nil {
		return nil, err
	}

	// Content
	var contentIds = gutil.ListItemValuesUnique(out.List, "Reply", "TargetId")
	err = dao.Content.Ctx(ctx).Where(dao.Content.Columns().Id, contentIds).
		ScanList(&out.List, "Content", "Reply", "id:TargetId")
	if err != nil {
		return nil, err
	}

	// Category
	var categoryIds = gutil.ListItemValuesUnique(out.List, "Content", "CategoryId")
	err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().Id, categoryIds).
		ScanList(&out.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}

	return out, nil
}
