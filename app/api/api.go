package api

type CommonListReq struct {
	Page int `json:"page" description:"分页号码，默认1"  default:"1" v:"min:0#分页号码错误"`
	Size int `json:"size" description:"分页数量，最大50" default:"10" v:"max:50#分页数量最大50条"`
}
