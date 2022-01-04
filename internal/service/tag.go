package service

type sTag struct{}

var insTag = sTag{}

// 标签管理服务
func Tag() *sTag {
	return &insTag
}
