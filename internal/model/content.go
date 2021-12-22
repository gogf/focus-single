package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"focus-single/internal/model/entity"
)

// ContentGetListInput 获取内容列表
type ContentGetListInput struct {
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	UserId     uint   // 要查询的用户ID
}

// ContentGetListOutput 查询列表结果
type ContentGetListOutput struct {
	List  []ContentGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

// ContentSearchInput 搜索列表
type ContentSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// ContentSearchOutput 搜索列表结果
type ContentSearchOutput struct {
	List  []ContentSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int            `json:"stats"` // 搜索统计
	Page  int                       `json:"page"`  // 分页码
	Size  int                       `json:"size"`  // 分页数量
	Total int                       `json:"total"` // 数据总数
}

type ContentGetListOutputItem struct {
	Content  *ContentListItem         `json:"content"`
	Category *ContentListCategoryItem `json:"category"`
	User     *ContentListUserItem     `json:"user"`
}

type ContentSearchOutputItem struct {
	ContentGetListOutputItem
}

// ContentGetDetailOutput 查询详情结果
type ContentGetDetailOutput struct {
	Content *entity.Content `json:"content"`
	User    *entity.User    `json:"user"`
}

// ContentCreateUpdateBase 创建/修改内容基类
type ContentCreateUpdateBase struct {
	Type       string   // 内容模型
	CategoryId uint     // 栏目ID
	Title      string   // 标题
	Content    string   // 内容
	Brief      string   // 摘要
	Thumb      string   // 缩略图
	Tags       []string // 标签名称列表，以JSON存储
	Referer    string   // 内容来源，例如github/gitee
}

// ContentCreateInput 创建内容
type ContentCreateInput struct {
	ContentCreateUpdateBase
	UserId uint
}

// ContentCreateOutput 创建内容返回结果
type ContentCreateOutput struct {
	ContentId uint `json:"content_id"`
}

// ContentUpdateInput 修改内容
type ContentUpdateInput struct {
	ContentCreateUpdateBase
	Id uint
}

// ContentListItem 主要用于列表展示
type ContentListItem struct {
	Id         uint        `json:"id"`          // 自增ID
	CategoryId uint        `json:"category_id"` // 栏目ID
	UserId     uint        `json:"user_id"`     // 用户ID
	Title      string      `json:"title"`       // 标题
	Sort       uint        `json:"sort"`        // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief      string      `json:"brief"`       // 摘要
	Thumb      string      `json:"thumb"`       // 缩略图
	Tags       string      `json:"tags"`        // 标签名称列表，以JSON存储
	Referer    string      `json:"referer"`     // 内容来源，例如github/gitee
	Status     uint        `json:"status"`      // 状态 0: 正常, 1: 禁用
	ViewCount  uint        `json:"view_count"`  // 浏览数量
	ReplyCount uint        `json:"reply_count"` // 回复数量
	ZanCount   uint        `json:"zan_count"`   // 赞
	CaiCount   uint        `json:"cai_count"`   // 踩
	CreatedAt  *gtime.Time `json:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"`  // 修改时间
}

// ContentListCategoryItem 绑定到Content列表中的栏目信息
type ContentListCategoryItem struct {
	Id          uint   `json:"id"`           // 分类ID，自增主键
	Name        string `json:"name"`         // 分类名称
	Thumb       string `json:"thumb"`        // 封面图
	ContentType string `json:"content_type"` // 内容类型：content, ask, article, reply

}

// ContentListUserItem 绑定到Content列表中的用户信息
type ContentListUserItem struct {
	Id       uint   `json:"id"`       // UID
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像地址
}

// ContentDetail Content详情
type ContentDetail struct {
	Content entity.Content
	User    entity.User
}
