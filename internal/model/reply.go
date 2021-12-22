package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReplyCreateInput 创建内容
type ReplyCreateInput struct {
	Title      string
	ParentId   uint   // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string // 评论类型: topic, ask, article, reply
	TargetId   uint   // 对应内容ID
	Content    string // 回复内容
	UserId     uint
}

// ReplyGetListInput 查询回复列表
type ReplyGetListInput struct {
	Page       int    // 分页码
	Size       int    // 分页数量
	TargetType string // 数据类型
	TargetId   uint   // 数据ID
	UserId     uint   // 用户ID
}

// ReplyGetListOutput 查询列表结果
type ReplyGetListOutput struct {
	List  []ReplyGetListOutputItem `json:"list"`  // 列表
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
}

// ReplyGetListOutputItem 查询列表结果项
type ReplyGetListOutputItem struct {
	Reply    *ReplyListItem           `json:"reply"`
	User     *ReplyListUserItem       `json:"user"`
	Content  *ContentListItem         `json:"content"`
	Category *ContentListCategoryItem `json:"category"`
}

// ReplyListItem 评论列表项
type ReplyListItem struct {
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
type ReplyListUserItem struct {
	Id       uint   `json:"id"`       // UID
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像地址
}
