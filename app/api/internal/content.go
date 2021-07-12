package internal

import "focus/app/model"

// 获取内容列表
type ContentGetListReq struct {
	model.ContentGetListInput
	CategoryId uint `p:"cate"`                    // 栏目ID
	Page       int  `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Size       int  `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
}

// 查看内容详情
type ContentDetailReq struct {
	Id uint `v:"min:1#请选择查看的内容"`
}

// 展示创建内容页面
type ContentCreateReq struct {
	Type string `v:"required#请选择需要创建的内容类型"`
}

// 展示修改内容页面
type ContentUpdateReq struct {
	Id uint `v:"min:1#请选择需要修改的内容"`
}

// 执行创建内容
type ContentDoCreateReq struct {
	model.ContentCreateInput
}

// 执行修改内容
type ContentDoUpdateReq struct {
	model.ContentUpdateInput
	Id uint `v:"min:1#请选择需要修改的内容"` // 修改时ID不能为空
}

// 执行删除内容
type ContentDoDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的内容"` // 删除时ID不能为空
}

// 执行采纳回复
type ContentAdoptReplyReq struct {
	Id      uint `v:"min:1#请选择需要采纳回复的内容"` // 采纳回复时ID不能为空
	ReplyId uint `v:"min:1#请选择需要采纳的回复"`   // 采纳回复时回复ID不能为空
}

// 搜索列表
type ContentSearchReq struct {
	model.ContentSearchInput
	CategoryId uint `p:"cate"`                    // 栏目ID
	Page       int  `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Size       int  `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
}
