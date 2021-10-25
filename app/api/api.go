package api

type CommonListReq struct {
	Page int `json:"page" dc:"分页号码，默认1"  in:"query" default:"1" v:"min:0#分页号码错误"`
	Size int `json:"size" dc:"分页数量，最大50" in:"query" default:"10" v:"max:50#分页数量最大50条"`
}
