// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import "context"

type IInteract interface {
	Zan(ctx context.Context, targetType string, targetId uint) error
	CancelZan(ctx context.Context, targetType string, targetId uint) error
	DidIZan(ctx context.Context, targetType string, targetId uint) (bool, error)
	Cai(ctx context.Context, targetType string, targetId uint) error
	CancelCai(ctx context.Context, targetType string, targetId uint) error
	DidICai(ctx context.Context, targetType string, targetId uint) (bool, error)
}

var localInteract IInteract

func Interact() IInteract {
	if localInteract == nil {
		panic("implement not found for interface IInteract, forgot register?")
	}
	return localInteract
}

func RegisterInteract(i IInteract) {
	localInteract = i
}
