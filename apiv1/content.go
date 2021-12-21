package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ContentGetListCommonReq struct {
	Type       string `json:"type"   in:"query" dc:"内容模型"`
	CategoryId uint   `json:"cate"   in:"query" dc:"栏目ID"`
	Sort       int    `json:"sort"   in:"query" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`
	CommonPaginationReq
}
type ContentGetListCommonRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ContentShowCreateReq struct {
	g.Meta `path:"/content/create" method:"get" tags:"内容" summary:"展示创建内容页面"`
	Type   string `json:"type" v:"required#请选择需要创建的内容类型"`
}
type ContentShowCreateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ContentCreateReq struct {
	g.Meta     `path:"/content/create" method:"post" tags:"内容" summary:"创建内容接口"`
	Type       string   `json:"type"    v:"required#内容模型不能为空" dc:"内容模型"`
	CategoryId uint     `json:"cate"    v:"required#请选择栏目"      dc:"栏目ID"`
	Title      string   `json:"title"   v:"required#请输入标题"      dc:"标题"`
	Content    string   `json:"content" v:"required#请输入内容"      dc:"内容"`
	Brief      string   `json:"brief"   dc:"摘要"`
	Thumb      string   `json:"thumb"   dc:"缩略图"`
	Tags       []string `json:"tags"    dc:"标签名称列表，以JSON存储"`
	Referer    string   `json:"referer" dc:"内容来源，例如github/gitee"`
}
type ContentCreateRes struct {
	ContentId uint `json:"contentId"`
}

type ContentShowUpdateReq struct {
	g.Meta `path:"/content/update/{Id}" method:"get" tags:"内容" summary:"展示内容修改页面"`
	Id     uint `json:"id" dc:"内容id" v:"min:1#请选择需要修改的内容"`
}
type ContentShowUpdateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ContentUpdateReq struct {
	g.Meta     `path:"/content/update/{Id}" method:"post" tags:"内容" summary:"修改内容接口"`
	Id         uint     `json:"id"      v:"min:1#请选择需要修改的内容" dc:"内容Id"`
	Type       string   `json:"type"    v:"required#内容模型不能为空" dc:"内容模型"`
	CategoryId uint     `json:"cate"    v:"required#请选择栏目"      dc:"栏目ID"`
	Title      string   `json:"title"   v:"required#请输入标题"      dc:"标题"`
	Content    string   `json:"content" v:"required#请输入内容"      dc:"内容"`
	Brief      string   `json:"brief"   dc:"摘要"`
	Thumb      string   `json:"thumb"   dc:"缩略图"`
	Tags       []string `json:"tags"    dc:"标签名称列表，以JSON存储"`
	Referer    string   `json:"referer" dc:"内容来源，例如github/gitee"`
}
type ContentUpdateRes struct{}

type ContentDeleteReq struct {
	g.Meta `path:"/content/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type ContentDeleteRes struct{}
