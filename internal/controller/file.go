package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/service/file"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 文件管理
var File = cFile{}

type cFile struct{}

func (a *cFile) Upload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	result, err := file.Upload(ctx, file.UploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.FileUploadRes{
		Name: result.Name,
		Url:  result.Url,
	}
	return
}
