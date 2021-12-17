package utils

import (
	"github.com/russross/blackfriday/v2"
)

// MarkdownToHtml 解析markdown为html
func MarkdownToHtml(mdContent string) string {
	return string(blackfriday.Run([]byte(mdContent)))
}
