package internal

// 上传文件返回结果
type FileUploadRes struct {
	Name string // 文件名称
	Url  string // 访问URL，可能只是URI
}
