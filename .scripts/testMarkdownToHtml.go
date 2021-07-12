package main

import (
	"fmt"
	"focus/library/utils"
	"github.com/gogf/gf/os/gfile"
)

func main() {
	content := gfile.GetContents("/Users/john/Workspace/Go/GOPATH/src/gitee.com/johng/focus/.scripts/test.md")

	fmt.Println(utils.MarkdownToHtml(content))

}
