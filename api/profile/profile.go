// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package profile

import (
	"context"

	"focus-single/api/profile/v1"
)

type IProfileV1 interface {
	ProfileIndex(ctx context.Context, req *v1.ProfileIndexReq) (res *v1.ProfileIndexRes, err error)
	ProfileUpdate(ctx context.Context, req *v1.ProfileUpdateReq) (res *v1.ProfileUpdateRes, err error)
	ProfileAvatar(ctx context.Context, req *v1.ProfileAvatarReq) (res *v1.ProfileAvatarRes, err error)
	ProfileUpdateAvatar(ctx context.Context, req *v1.ProfileUpdateAvatarReq) (res *v1.ProfileUpdateAvatarRes, err error)
	ProfilePassword(ctx context.Context, req *v1.ProfilePasswordReq) (res *v1.ProfilePasswordRes, err error)
	ProfileUpdatePassword(ctx context.Context, req *v1.ProfileUpdatePasswordReq) (res *v1.ProfileUpdatePasswordRes, err error)
	ProfileMessage(ctx context.Context, req *v1.ProfileMessageReq) (res *v1.ProfileMessageRes, err error)
}
