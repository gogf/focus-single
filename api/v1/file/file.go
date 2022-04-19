package file

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta `path:"/file" method:"post" mime:"multipart/form-data" tags:"工具" summary:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}
type UploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url"  dc:"访问URL，可能只是URI"`
}
