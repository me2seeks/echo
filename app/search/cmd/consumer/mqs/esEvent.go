package mqs

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/me2seeks/echo-hub/app/search/cmd/consumer/internal/svc"
	"github.com/me2seeks/echo-hub/common/es"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type EsEvent struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEsEvent(ctx context.Context, svcCtx *svc.ServiceContext) *EsEvent {
	return &EsEvent{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EsEvent) Consume(ctx context.Context, key, val string) error {
	logx.Infof("EsEvent key :%s , val :%s", key, val)
	var event kqueue.EsEvent
	err := json.Unmarshal(tool.StringToBytes(val), &event)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.UnmarshalError), "unmarshal event err:%v", err)
	}

	switch event.Type {
	case kqueue.Register:
		userJSON, err := json.Marshal(es.User{
			ID:       event.ID,
			Nickname: event.Nickname,
			Handle:   event.Content,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "marshal user err:%v", err)
		}
		req := esapi.IndexRequest{
			Index:      "users",
			DocumentID: strconv.FormatInt(event.ID, 10),
			Body:       strings.NewReader(string(userJSON)),
			Refresh:    "true",
		}
		res, err := req.Do(ctx, l.svcCtx.EsClient)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.EsError), "es request err:%v", err)
		}
		defer res.Body.Close()
		if res.IsError() {
			return errors.Wrapf(xerr.NewErrCode(xerr.EsError), "es response err:%v", res.String())
		}
	case kqueue.Feed:
		userID, err := strconv.ParseInt(event.Nickname, 10, 64)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.Str2Int64Error), "invalid event err:%v", err)
		}
		feedJSON, err := json.Marshal(es.Feed{
			ID:      event.ID,
			UserID:  userID,
			Content: event.Content,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "marshal feed err:%v", err)
		}
		req := esapi.IndexRequest{
			Index:      "feeds",
			DocumentID: strconv.FormatInt(event.ID, 10),
			Body:       strings.NewReader(string(feedJSON)),
			Refresh:    "true",
		}
		res, err := req.Do(ctx, l.svcCtx.EsClient)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.EsError), "es request err:%v", err)
		}
		defer res.Body.Close()
		if res.IsError() {
			return errors.Wrapf(xerr.NewErrCode(xerr.EsError), "es response err:%v", res.String())
		}
	default:
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidEvent), "invalid event type:%d", event.Type)
	}

	return nil
}
