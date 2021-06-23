package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// 交互管理服务
var Interact = interactService{}

type interactService struct{}

const (
	contextMapKeyForMyInteractList = "ContextMapKeyForMyInteractList"
)

// 赞
func (s *interactService) Zan(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := shared.Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Data(&model.Interact{
			UserId:     customCtx.User.Id,
			TargetId:   targetId,
			TargetType: targetType,
			Type:       model.InteractTypeZan,
		}).FieldsEx(dao.Interact.C.Id).InsertIgnore()
		if err != nil {
			return err
		}

		if n, _ := r.RowsAffected(); n == 0 {
			return gerror.New("您已经赞过啦")
		}
		return s.updateCount(ctx, model.InteractTypeZan, targetType, targetId, 1)
	})
}

// 取消赞
func (s *interactService) CancelZan(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := shared.Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Where(g.Slice{
			dao.Interact.C.UserId, shared.Context.Get(ctx).User.Id,
			dao.Interact.C.TargetId, targetId,
			dao.Interact.C.TargetType, targetType,
			dao.Interact.C.Type, model.InteractTypeZan,
		}).Delete()
		if err != nil {
			return err
		}
		if n, _ := r.RowsAffected(); n == 0 {
			return nil
		}
		return s.updateCount(ctx, model.InteractTypeZan, targetType, targetId, -1)
	})
}

// 我是否有对指定内容赞
func (s *interactService) DidIZan(ctx context.Context, targetType string, targetId uint) (bool, error) {
	list, err := s.getMyList(ctx)
	if err != nil {
		return false, err
	}
	for _, v := range list {
		if v.TargetId == targetId && v.TargetType == targetType && v.Type == model.InteractTypeZan {
			return true, nil
		}
	}
	return false, nil
}

// 踩
func (s *interactService) Cai(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := shared.Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Data(&model.Interact{
			UserId:     customCtx.User.Id,
			TargetId:   targetId,
			TargetType: targetType,
			Type:       model.InteractTypeCai,
		}).FieldsEx(dao.Interact.C.Id).InsertIgnore()
		if err != nil {
			return err
		}
		if n, _ := r.RowsAffected(); n == 0 {
			return gerror.New("您已经踩过啦")
		}
		return s.updateCount(ctx, model.InteractTypeCai, targetType, targetId, 1)
	})
}

// 取消踩
func (s *interactService) CancelCai(ctx context.Context, targetType string, targetId uint) error {
	return dao.Interact.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		customCtx := shared.Context.Get(ctx)
		if customCtx == nil || customCtx.User == nil {
			return nil
		}
		r, err := dao.Interact.Where(g.Slice{
			dao.Interact.C.UserId, shared.Context.Get(ctx).User.Id,
			dao.Interact.C.TargetId, targetId,
			dao.Interact.C.TargetType, targetType,
			dao.Interact.C.Type, model.InteractTypeCai,
		}).Delete()
		if err != nil {
			return err
		}
		if n, _ := r.RowsAffected(); n == 0 {
			return nil
		}
		return s.updateCount(ctx, model.InteractTypeCai, targetType, targetId, -1)
	})
}

// 我是否有对指定内容踩
func (s *interactService) DidICai(ctx context.Context, targetType string, targetId uint) (bool, error) {
	list, err := s.getMyList(ctx)
	if err != nil {
		return false, err
	}
	for _, v := range list {
		if v.TargetId == targetId && v.TargetType == targetType && v.Type == model.InteractTypeCai {
			return true, nil
		}
	}
	return false, nil
}

// 获得我的互动数据列表，内部带请求上下文缓存
func (s *interactService) getMyList(ctx context.Context) ([]*model.Interact, error) {
	customCtx := shared.Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil, nil
	}
	if v, ok := customCtx.Data[contextMapKeyForMyInteractList]; ok {
		return v.([]*model.Interact), nil
	}
	var list []*model.Interact
	err := dao.Interact.Where(dao.Interact.C.UserId, customCtx.User.Id).Scan(&list)
	if err != nil {
		return nil, err
	}
	customCtx.Data[contextMapKeyForMyInteractList] = list
	return list, err
}

func (s *interactService) updateCount(ctx context.Context, interactType int, targetType string, targetId uint, count int) error {
	defer func() {
		// 清空上下文对应的互动数据缓存
		if customCtx := shared.Context.Get(ctx); customCtx != nil {
			delete(customCtx.Data, contextMapKeyForMyInteractList)
		}
	}()
	var err error
	switch targetType {
	// 内容赞踩
	case model.InteractTargetTypeContent:
		switch interactType {
		case model.InteractTypeZan:
			_, err = dao.Content.Ctx(ctx).
				Where(dao.Content.C.Id, targetId).
				WhereGTE(dao.Content.C.ZanCount, 0).
				Increment(dao.Content.C.ZanCount, count)
			if err != nil {
				return err
			}

		case model.InteractTypeCai:
			_, err = dao.Content.Ctx(ctx).
				Where(dao.Content.C.Id, targetId).
				WhereGTE(dao.Content.C.CaiCount, 0).
				Increment(dao.Content.C.CaiCount, count)
			if err != nil {
				return err
			}
		}
	// 评论赞踩
	case model.InteractTargetTypeReply:
		switch interactType {
		case model.InteractTypeZan:
			_, err = dao.Reply.Ctx(ctx).
				Where(dao.Content.C.Id, targetId).
				WhereGTE(dao.Content.C.ZanCount, 0).
				Increment(dao.Content.C.ZanCount, count)
			if err != nil {
				return err
			}

		case model.InteractTypeCai:
			_, err = dao.Reply.Ctx(ctx).
				Where(dao.Content.C.Id, targetId).
				WhereGTE(dao.Content.C.CaiCount, 0).
				Increment(dao.Content.C.CaiCount, count)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
