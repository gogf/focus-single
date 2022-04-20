package reply

import (
	"context"

	"focus-single/internal/model/do"
	"focus-single/internal/service/bizctx"
	"focus-single/internal/service/content"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gutil"

	"focus-single/internal/dao"
	"focus-single/internal/model/entity"
)

// ReplyCreateInput 创建内容
type CreateInput struct {
	Title      string
	ParentId   uint   // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string // 评论类型: topic, ask, article, reply
	TargetId   uint   // 对应内容ID
	Content    string // 回复内容
	UserId     uint
}

// ReplyGetListInput 查询回复列表
type GetListInput struct {
	Page       int    // 分页码
	Size       int    // 分页数量
	TargetType string // 数据类型
	TargetId   uint   // 数据ID
	UserId     uint   // 用户ID
}

// ReplyGetListOutput 查询列表结果
type GetListOutput struct {
	List  []GetListOutputItem `json:"list"`  // 列表
	Page  int                 `json:"page"`  // 分页码
	Size  int                 `json:"size"`  // 分页数量
	Total int                 `json:"total"` // 数据总数
}

// ReplyGetListOutputItem 查询列表结果项
type GetListOutputItem struct {
	Reply    *ListItem                 `json:"reply"`
	User     *ListUserItem             `json:"user"`
	Content  *content.ListItem         `json:"content"`
	Category *content.ListCategoryItem `json:"category"`
}

// ReplyListItem 评论列表项
type ListItem struct {
	Id         uint        `json:"id"`          // 回复ID
	ParentId   uint        `json:"parent_id"`   // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string      `json:"target_type"` // 评论类型: topic, ask, article, reply
	TargetId   uint        `json:"target_id"`   // 对应内容ID
	UserId     uint        `json:"user_id"`     // 网站用户ID
	ZanCount   uint        `json:"zan_count"`   // 赞
	CaiCount   uint        `json:"cai_count"`   // 踩
	Title      string      `json:"title"`       // 回复标题
	Content    string      `json:"content"`     // 回复内容
	CreatedAt  *gtime.Time `json:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"`  //
}

// ReplyListUserItem 绑定到Content列表中的用户信息
type ListUserItem struct {
	Id       uint   `json:"id"`       // UID
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像地址
}

// Create 创建回复
func Create(ctx context.Context, in CreateInput) error {
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
func GetList(ctx context.Context, in GetListInput) (out *GetListOutput, err error) {
	out = &GetListOutput{
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
