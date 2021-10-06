package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 上传文件
type FileUploadReq struct {
	g.Meta `method:"post" mime:"multipart/form-data" summary:"上传文件" tags:"工具"`
	File   *ghttp.UploadFile `json:"file" description:"选择上传文件"`
}
type FileUploadRes struct {
	Name string `json:"name" description:"文件名称"`
	Url  string `json:"url" description:"访问URL，可能只是URI"`
}
