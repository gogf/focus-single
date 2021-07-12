package internal

// 赞
type InteractZanReq struct {
	Id   uint   `v:"min:1#请选择需要赞的内容"`
	Type string `v:"required#请提交需要赞的内容类型"` // content, reply
}

// 取消赞
type InteractCancelZanReq struct {
	Id   uint   `v:"min:1#请选择需要取消赞的内容"`
	Type string `v:"required#请提交需要取消赞的内容类型"` // content, reply
}

// 踩
type InteractCaiReq struct {
	Id   uint   `v:"min:1#请选择需要踩的内容"`
	Type string `v:"required#请提交需要踩的内容类型"` // content, reply
}

// 取消踩
type InteractCancelCaiReq struct {
	Id   uint   `v:"min:1#请选择需要取消踩的内容"`
	Type string `v:"required#请提交需要取消踩的内容类型"` // content, reply
}
