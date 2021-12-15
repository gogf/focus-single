package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 上传文件
type FileUploadReq struct {
	g.Meta `method:"post" mime:"multipart/form-data" summary:"上传文件" tags:"工具"`
	File   *ghttp.UploadFile `json:"file" dc:"选择上传文件"`
}
type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url"  dc:"访问URL，可能只是URI"`
}
