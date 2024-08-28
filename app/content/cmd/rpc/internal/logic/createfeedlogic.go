package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/uniqueid"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFeedLogic {
	return &CreateFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFeedLogic) CreateFeed(in *pb.CreateFeedReq) (*pb.CreateFeedResp, error) {
	id := uniqueid.GenFeedID()
	_, err := l.svcCtx.FeedsModel.Insert(l.ctx, nil, &model.Feeds{
		Id:      id,
		Content: in.Content,
		UserId:  in.UserID,
		Media0:  sql.NullString{String: in.Media0, Valid: in.Media0 != ""},
		Media1:  sql.NullString{String: in.Media1, Valid: in.Media1 != ""},
		Media2:  sql.NullString{String: in.Media2, Valid: in.Media2 != ""},
		Media3:  sql.NullString{String: in.Media3, Valid: in.Media3 != ""},
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "CreateFeed insert feed failed: %v", err)
	}

	go func() {
		msg := kqueue.EsEvent{
			Type:      kqueue.Feed,
			ID:        id,
			UserID:    in.UserID,
			Content:   in.Content,
			CreatedAt: time.Now(),
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			logx.Errorf("CreateFeed Marshal EsEvent failed  Type:%d,ID:%d,UserID:%d,Content:%s,CreateAt:%v,err:%v", kqueue.Feed, id, in.UserID, in.Content, time.Now(), err)
			return
		}
		err = l.svcCtx.KqPusherEsEventClient.Push(l.ctx, tool.BytesToString(msgBytes))
		if err != nil {
			logx.Errorf("CreateFeed Push EsEvent failed  Type:%d,ID:%d,UserID:%d,Content:%s,CreateAt:%v,err:%v", kqueue.Feed, id, in.UserID, in.Content, time.Now(), err)
			return
		}
	}()

	go func() {
		msg := kqueue.CountEvent{
			Type:      kqueue.Feed,
			SourceID:  in.UserID,
			IsComment: false,
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			logx.Errorf("CreateFeed Marshal CountEvent failed  Type:%d,SourceID:%d,IsComment:%t,err:%v", kqueue.Feed, in.UserID, false, err)
			return
		}
		userIDStr := strconv.FormatInt(in.UserID, 10)
		err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, userIDStr, tool.BytesToString(msgBytes))
		if err != nil {
			logx.Errorf("CreateFeed PushWithKey CountEvent failed  Type:%d,SourceID:%d,IsComment:%t,err:%v", kqueue.Feed, in.UserID, false, err)
			return
		}
	}()

	return &pb.CreateFeedResp{
		Id: id,
	}, nil
}
