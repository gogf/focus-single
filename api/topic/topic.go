// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package topic

import (
	"context"

	"focus-single/api/topic/v1"
)

type ITopicV1 interface {
	TopicIndex(ctx context.Context, req *v1.TopicIndexReq) (res *v1.TopicIndexRes, err error)
	TopicDetail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error)
}
