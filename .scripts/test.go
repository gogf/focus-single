package main

import (
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Dump(g.Cfg("admin").Get("."))
	//s, _ := gparser.VarToTomlString(g.Map{
	//	"menus": model.MenuAdmin,
	//})
	//fmt.Println(s)
}
