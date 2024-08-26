package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedsByIDByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedsByIDByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedsByIDByPageLogic {
	return &GetFeedsByIDByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFeedsByIDByPageLogic) GetFeedsByIDByPage(in *pb.GetFeedsByIDByPageReq) (*pb.GetFeedsByIDByPageResp, error) {
	feeds, total, err := l.svcCtx.FeedsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.FeedsModel.SelectBuilder().Columns("id, user_id, content, media0, media1, media2, media3, create_at").Where("id in (?,?,...) ", in.FeedID), in.Page, in.PageSize, "id DESC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetFeedsByIDByPage FindPageListByPageWithTotal err:%v", err)
	}
	var resp pb.GetFeedsByIDByPageResp
	for _, feed := range feeds {
		resp.Feeds = append(resp.Feeds, &pb.Feed{
			Id:         feed.Id,
			UserID:     feed.UserId,
			Content:    feed.Content,
			Media0:     feed.Media0.String,
			Media1:     feed.Media1.String,
			Media2:     feed.Media2.String,
			Media3:     feed.Media3.String,
			CreateTime: timestamppb.New(feed.CreateAt),
		})
	}
	resp.Total = total

	return &resp, nil
}
