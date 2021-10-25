package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ContentListCommonReq struct {
	CommonListReq
	CategoryId uint `json:"cate" required:"true" dc:"栏目ID"`
}

// 获取内容列表
type ContentGetListReq struct {
	g.Meta `method:"get" summary:"展示内容列表页面" tags:"内容"`
	ContentListCommonReq
	Type       string `json:"type"   in:"query" dc:"内容模型"`
	CategoryId uint   `json:"cate"   in:"query" dc:"栏目ID"`
	Page       int    `json:"page"   in:"query" dc:"分页号码"`
	Size       int    `json:"size"   in:"query" dc:"分页数量，最大50"`
	Sort       int    `json:"sort"   in:"query" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`
	UserId     uint   `json:"userId" in:"query" dc:"要查询的用户ID"`
}
type ContentGetListRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 查看内容详情
type ContentDetailReq struct {
	g.Meta `method:"get" summary:"展示内容详情页面" tags:"内容"`
	Id     uint `json:"id" in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type ContentDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 展示创建内容页面
type ContentCreateReq struct {
	g.Meta `method:"get" summary:"展示创建内容页面" tags:"内容"`
	Type   string `json:"type" v:"required#请选择需要创建的内容类型"`
}
type ContentCreateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 展示修改内容页面
type ContentUpdateReq struct {
	g.Meta `method:"get" summary:"展示内容修改页面" tags:"内容"`
	Id     uint `json:"id" in:"path" dc:"内容id" v:"min:1#请选择需要修改的内容"`
}
type ContentUpdateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// 执行创建内容
type ContentDoCreateReq struct {
	g.Meta     `method:"post" summary:"创建内容接口" tags:"内容"`
	Type       string   `json:"type"    v:"required|内容模型不能为空" dc:"内容模型"`
	CategoryId uint     `json:"cate"    v:"required|请选择栏目"      dc:"栏目ID"`
	Title      string   `json:"title"   v:"required|请输入标题"      dc:"标题"`
	Content    string   `json:"content" v:"required|请输入内容"      dc:"内容"`
	Brief      string   `json:"brief"   dc:"摘要"`
	Thumb      string   `json:"thumb"   dc:"缩略图"`
	Tags       []string `json:"tags"    dc:"标签名称列表，以JSON存储"`
	Referer    string   `json:"referer" dc:"内容来源，例如github/gitee"`
}
type ContentDoCreateRes struct {
	ContentId uint `json:"contentId"`
}

// 执行修改内容
type ContentDoUpdateReq struct {
	g.Meta     `method:"post" summary:"修改内容接口" tags:"内容"`
	Id         uint     `json:"id" in:"path" v:"min:1#请选择需要修改的内容" dc:"内容Id"`
	Type       string   `json:"type"    v:"required|内容模型不能为空" dc:"内容模型"`
	CategoryId uint     `json:"cate"    v:"required|请选择栏目"      dc:"栏目ID"`
	Title      string   `json:"title"   v:"required|请输入标题"      dc:"标题"`
	Content    string   `json:"content" v:"required|请输入内容"      dc:"内容"`
	Brief      string   `json:"brief"   dc:"摘要"`
	Thumb      string   `json:"thumb"   dc:"缩略图"`
	Tags       []string `json:"tags"    dc:"标签名称列表，以JSON存储"`
	Referer    string   `json:"referer" dc:"内容来源，例如github/gitee"`
}
type ContentDoUpdateRes struct{}

// 执行删除内容
type ContentDoDeleteReq struct {
	g.Meta `method:"post" summary:"删除内容接口" tags:"内容"`
	Id     uint `json:"id" in:"path" dc:"内容id" v:"min:1#请选择需要删除的内容"`
}
type ContentDoDeleteRes struct{}

// 执行采纳回复
type ContentAdoptReplyReq struct {
	g.Meta  `method:"post" summary:"采纳指定回复作为内容答案" tags:"内容"`
	Id      uint `json:"id"      dc:"内容id" v:"min:1#请选择需要采纳回复的内容"`
	ReplyId uint `json:"replyId" dc:"回复id" v:"min:1#请选择需要采纳的回复"`
}
type ContentAdoptReplyRes struct{}

// 搜索列表
type ContentSearchReq struct {
	g.Meta `method:"get" summary:"展示内容搜索页面" tags:"内容"`
	ContentListCommonReq
	Key        string `json:"key" v:"required|请输入搜索关键字" dc:"关键字"`
	Type       string `json:"type" dc:"内容模型"`
	CategoryId uint   `json:"cate" dc:"栏目ID"`
	Page       int    `json:"page" dc:"分页号码"`
	Size       int    `json:"size" v:"max:50|最大分页数为50" dc:"分页数量，最大50"`
	Sort       int    `json:"sort" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`
}
type ContentSearchRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
