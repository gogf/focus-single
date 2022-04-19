package file

import (
	"context"

	v1 "focus-single/api/v1/file"
	"focus-single/internal/model"
	"focus-single/internal/service/file"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	result, err := file.Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.UploadRes{
		Name: result.Name,
		Url:  result.Url,
	}
	return
}
