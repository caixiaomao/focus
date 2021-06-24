package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/index/internal/define"
	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

// 评论/回复管理服务
var Reply = replyService{}

type replyService struct{}

// 创建回复
func (s *replyService) Create(ctx context.Context, input define.ReplyCreateInput) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 覆盖用户ID
		input.UserId = shared.Context.Get(ctx).User.Id
		_, err := dao.Reply.Ctx(ctx).Data(input).Insert()
		if err == nil {
			err = Content.AddReplyCount(ctx, input.TargetId, 1)
		}
		return err
	})
}

// 删除回复(硬删除)
func (s *replyService) Delete(ctx context.Context, id uint) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var reply *model.Reply
		err := dao.Reply.Ctx(ctx).WherePri(id).Scan(&reply)
		if err != nil {
			return err
		}
		// 删除回复记录
		_, err = dao.Reply.Ctx(ctx).Where(g.Map{
			dao.Reply.C.Id:     id,
			dao.Reply.C.UserId: shared.Context.Get(ctx).User.Id,
		}).Delete()
		if err == nil {
			// 回复统计-1
			err = Content.AddReplyCount(ctx, reply.TargetId, -1)
			if err != nil {
				return err
			}
			// 判断回复是否采纳
			var content *model.Content
			err = dao.Content.Ctx(ctx).WherePri(reply.TargetId).Scan(&content)
			if err == nil && content != nil && content.AdoptedReplyId == id {
				err = Content.UnacceptedReply(ctx, reply.TargetId)
			}
		}
		return err
	})
}

// 删除回复(硬删除)
func (s *replyService) DeleteByUserContentId(ctx context.Context, userId, contentId uint) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除内容对应的回复
		_, err := dao.Reply.Ctx(ctx).Where(g.Map{
			dao.Reply.C.TargetId: contentId,
			dao.Reply.C.UserId:   userId,
		}).Delete()
		return err
	})
}

// 获取回复列表
func (s *replyService) GetList(ctx context.Context, input define.ReplyGetListInput) (output *define.ReplyGetListOutput, err error) {
	output = &define.ReplyGetListOutput{
		Page: input.Page,
		Size: input.Size,
	}
	m := dao.Reply.Ctx(ctx).Fields(model.ReplyListItem{})
	if input.TargetType != "" {
		m = m.Where(dao.Reply.C.TargetType, input.TargetType)
	}
	if input.TargetId > 0 {
		m = m.Where(dao.Reply.C.TargetId, input.TargetId)
	}
	if input.UserId > 0 {
		m = m.Where(dao.Reply.C.UserId, input.UserId)
	}

	err = m.Page(input.Page, input.Size).OrderDesc(dao.Content.C.Id).ScanList(&output.List, "Reply")
	if err != nil {
		return nil, err
	}
	if len(output.List) == 0 {
		return nil, nil
	}
	// User
	if err = m.ScanList(&output.List, "Reply"); err != nil {
		return nil, err
	}
	err = dao.User.
		Fields(model.ReplyListUserItem{}).
		Where(dao.User.C.Id, gutil.ListItemValuesUnique(output.List, "Reply", "UserId")).
		ScanList(&output.List, "User", "Reply", "id:UserId")
	if err != nil {
		return nil, err
	}

	// Content
	err = dao.Content.Fields(dao.Content.C.Id, dao.Content.C.Title, dao.Content.C.CategoryId).
		Where(dao.Content.C.Id, gutil.ListItemValuesUnique(output.List, "Reply", "TargetId")).
		ScanList(&output.List, "Content", "Reply", "id:TargetId")
	if err != nil {
		return nil, err
	}

	// Category
	err = dao.Category.
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.C.Id, gutil.ListItemValuesUnique(output.List, "Content", "CategoryId")).
		ScanList(&output.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}

	return output, nil
}
