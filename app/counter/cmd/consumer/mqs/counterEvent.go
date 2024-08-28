package mqs

import (
	"context"
	"encoding/json"

	"github.com/me2seeks/echo-hub/app/counter/cmd/consumer/internal/svc"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CounterEvent struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCounterEvent(ctx context.Context, svcCtx *svc.ServiceContext) *CounterEvent {
	return &CounterEvent{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CounterEvent) Consume(ctx context.Context, key, val string) error {
	logx.Debugf("CounterEvent key :%v , val :%v", key, val)
	var event kqueue.CountEvent
	err := json.Unmarshal(tool.StringToBytes(val), &event)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.UnmarshalError), "unmarshal event err:%v", err)
	}

	// TODO 多条消息一起消费
	switch event.Type {
	case kqueue.Follow:
		err := l.svcCtx.UserStateModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			err := l.svcCtx.UserStateModel.IncreaseFollowerCount(l.ctx, session, event.TargetID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase follower count  SourceID:%d,TargetID:%d,err:%v", event.SourceID, event.TargetID, err)
			}
			err = l.svcCtx.UserStateModel.IncreaseFollowingCount(l.ctx, session, event.SourceID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase following count SourceID:%d,TargetID:%d,err:%v", event.SourceID, event.TargetID, err)
			}
			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "follow event Trans SourceID:%d,TargetID:%d,err:%v", event.SourceID, event.TargetID, err)
		}
	case kqueue.UnFollow:
		err := l.svcCtx.UserStateModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			err := l.svcCtx.UserStateModel.DecreaseFollowerCount(l.ctx, session, event.TargetID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease follower count err:%v", err)
			}
			err = l.svcCtx.UserStateModel.DecreaseFollowingCount(l.ctx, session, event.SourceID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease following count err:%v", err)
			}
			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "unfollow eventTrans SourceID:%d,TargetID:%d,err:%v", event.SourceID, event.TargetID, err)
		}
	case kqueue.Like:
		if !event.IsComment {
			err := l.svcCtx.FeedCounterModel.IncreaseLikeCount(l.ctx, nil, event.TargetID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase like count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
			}
			return nil
		}
		err := l.svcCtx.CommentCounterModel.IncreaseLikeCount(l.ctx, nil, event.TargetID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase like count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
		}
	case kqueue.UnLike:
		if !event.IsComment {
			err := l.svcCtx.FeedCounterModel.DecreaseLikeCount(l.ctx, nil, event.TargetID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease like count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
			}
			return nil
		}
		err := l.svcCtx.CommentCounterModel.DecreaseLikeCount(l.ctx, nil, event.TargetID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease like count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
		}
	case kqueue.Comment:
		if !event.IsComment {
			err := l.svcCtx.FeedCounterModel.IncreaseCommentCount(l.ctx, nil, event.TargetID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase comment count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
			}
			return nil
		}
		err := l.svcCtx.CommentCounterModel.IncreaseCommentCount(l.ctx, nil, event.TargetID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase comment count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
		}
	case kqueue.UnComment:
		if !event.IsComment {
			err := l.svcCtx.FeedCounterModel.DecreaseCommentCount(l.ctx, nil, event.TargetID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease comment count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
			}
			return nil
		}
		err := l.svcCtx.CommentCounterModel.DecreaseCommentCount(l.ctx, nil, event.TargetID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease comment count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
		}
	case kqueue.View:
		if !event.IsComment {
			err := l.svcCtx.FeedCounterModel.IncreaseViewCount(l.ctx, nil, event.TargetID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase view count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
			}
			return nil
		}
		err := l.svcCtx.CommentCounterModel.IncreaseViewCount(l.ctx, nil, event.TargetID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase view count TargetID:%d,IsComment:%t,err:%v", event.TargetID, event.IsComment, err)
		}
	case kqueue.Feed:
		err := l.svcCtx.UserStateModel.IncreaseFeedCount(l.ctx, nil, event.SourceID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase feed count SourceID:%d,err:%v", event.SourceID, err)
		}
	case kqueue.DeleteFeed:
		err := l.svcCtx.UserStateModel.DecreaseFeedCount(l.ctx, nil, event.SourceID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease feed count SourceID:%d,err:%v", event.SourceID, err)
		}
	default:
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidEvent), "invalid event type:%d", event.Type)
	}

	return nil
}
