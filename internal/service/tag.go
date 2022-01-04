package service

// 标签管理服务
var insTag = sTag{}

type sTag struct{}

func Tag() *sTag {
	return &insTag
}
