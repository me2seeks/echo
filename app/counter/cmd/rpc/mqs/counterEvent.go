package mqs

import (
	"context"
	"encoding/json"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
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
	logx.Infof("CounterEvent key :%s , val :%s", key, val)
	var event kqueue.Event
	err := json.Unmarshal(tool.StringToBytes(val), &event)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.UnmarshalError), "unmarshal event err:%v", err)
	}

	// TODO 多条消息一起消费
	if event.IsComment {
		switch event.Type {
		case kqueue.Follow:
			err := l.svcCtx.UserStateModel.IncreaseFollowerCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase follower count err:%v", err)
			}
		case kqueue.UnFollow:
			err := l.svcCtx.UserStateModel.DecreaseFollowerCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease follower count err:%v", err)
			}
		case kqueue.Like:
			if !event.IsComment {
				err := l.svcCtx.FeedCounterModel.IncreaseLikeCount(l.ctx, nil, event.ID)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase like count err:%v", err)
				}
				return nil
			}
			err := l.svcCtx.CommentCounterModel.IncreaseLikeCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase like count err:%v", err)
			}
		case kqueue.UnLike:
			if !event.IsComment {
				err := l.svcCtx.FeedCounterModel.DecreaseLikeCount(l.ctx, nil, event.ID)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease like count err:%v", err)
				}
				return nil
			}
			err := l.svcCtx.CommentCounterModel.DecreaseLikeCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease like count err:%v", err)
			}
		case kqueue.Comment:
			if !event.IsComment {
				err := l.svcCtx.FeedCounterModel.IncreaseCommentCount(l.ctx, nil, event.ID)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase comment count err:%v", err)
				}
				return nil
			}
			err := l.svcCtx.CommentCounterModel.IncreaseCommentCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase comment count err:%v", err)
			}
		case kqueue.UnComment:
			if !event.IsComment {
				err := l.svcCtx.FeedCounterModel.DecreaseCommentCount(l.ctx, nil, event.ID)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease comment count err:%v", err)
				}
				return nil
			}
			err := l.svcCtx.CommentCounterModel.DecreaseCommentCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease comment count err:%v", err)
			}
		case kqueue.View:
			if !event.IsComment {
				err := l.svcCtx.FeedCounterModel.IncreaseViewCount(l.ctx, nil, event.ID)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase view count err:%v", err)
				}
				return nil
			}
			err := l.svcCtx.CommentCounterModel.IncreaseViewCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase view count err:%v", err)
			}
		case kqueue.Feed:
			err := l.svcCtx.UserStateModel.IncreaseFeedCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "increase feed count err:%v", err)
			}
		case kqueue.DeleteFeed:
			err := l.svcCtx.UserStateModel.DecreaseFeedCount(l.ctx, nil, event.ID)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "decrease feed count err:%v", err)
			}
		default:
			return errors.Wrapf(xerr.NewErrCode(xerr.InvalidEvent), "invalid event type:%d", event.Type)
		}
	} else {
		switch event.Type {
		}
	}
	return nil
}
