package api

import (
	"focus/app/internal/model"
	"github.com/gogf/gf/frame/g"
)

type ContentListCommonReq struct {
	CategoryId uint `json:"cate" required:"true" description:"栏目ID"`
	CommonListReq
}

// 获取内容列表
type ContentGetListReq struct {
	g.Meta `method:"get" summary:"展示内容列表页面" tags:"内容"`
	ContentListCommonReq
	model.ContentGetListInput
}
type ContentGetListRes struct {
	g.Meta `mime:"text/html"`
}

// 查看内容详情
type ContentDetailReq struct {
	g.Meta `method:"get" summary:"展示内容详情页面" tags:"内容"`
	Id     uint `json:"id" in:"path" description:"内容id" v:"min:1#请选择查看的内容"`
}
type ContentDetailRes struct {
	g.Meta `mime:"text/html"`
}

// 展示创建内容页面
type ContentCreateReq struct {
	g.Meta `method:"get" summary:"展示创建内容页面" tags:"内容"`
	Type   string `v:"required#请选择需要创建的内容类型"`
}
type ContentCreateRes struct {
	g.Meta `mime:"text/html"`
}

// 展示修改内容页面
type ContentUpdateReq struct {
	g.Meta `method:"get" summary:"展示内容修改页面" tags:"内容"`
	Id     uint `json:"id" in:"path" description:"内容id" v:"min:1#请选择需要修改的内容"`
}
type ContentUpdateRes struct {
	g.Meta `mime:"text/html"`
}

// 执行创建内容
type ContentDoCreateReq struct {
	g.Meta `method:"post" summary:"创建内容接口" tags:"内容"`
	model.ContentCreateInput
}
type ContentDoCreateRes struct {
	ContentId uint `json:"contentId"`
}

// 执行修改内容
type ContentDoUpdateReq struct {
	g.Meta `method:"post" summary:"修改内容接口" tags:"内容"`
	model.ContentUpdateInput
	Id uint `json:"id" in:"path" description:"内容id" v:"min:1#请选择需要修改的内容"`
}
type ContentDoUpdateRes struct{}

// 执行删除内容
type ContentDoDeleteReq struct {
	g.Meta `method:"post" summary:"删除内容接口" tags:"内容"`
	Id     uint `json:"id" in:"path" description:"内容id" v:"min:1#请选择需要删除的内容"`
}
type ContentDoDeleteRes struct{}

// 执行采纳回复
type ContentAdoptReplyReq struct {
	g.Meta  `method:"post" summary:"采纳指定回复作为内容答案" tags:"内容"`
	Id      uint `json:"id"      description:"内容id" v:"min:1#请选择需要采纳回复的内容"`
	ReplyId uint `json:"replyId" description:"回复id" v:"min:1#请选择需要采纳的回复"`
}
type ContentAdoptReplyRes struct{}

// 搜索列表
type ContentSearchReq struct {
	g.Meta `method:"get" summary:"展示内容搜索页面" tags:"内容"`
	model.ContentSearchInput
	ContentListCommonReq
}
type ContentSearchRes struct {
	g.Meta `mime:"text/html"`
}
