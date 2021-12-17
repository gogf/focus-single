package model

import "github.com/gogf/gf/v2/net/ghttp"

// FileUploadInput 上传文件输入参数
type FileUploadInput struct {
	File       *ghttp.UploadFile // 上传文件对象
	Name       string            // 自定义文件名称
	RandomName bool              // 是否随机命名文件
}

// FileUploadOutput 上传文件返回参数
type FileUploadOutput struct {
	Id   uint   // 数据表ID
	Name string // 文件名称
	Path string // 本地路径
	Url  string // 访问URL，可能只是URI
}
