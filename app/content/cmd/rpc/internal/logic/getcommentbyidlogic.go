package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByIDLogic {
	return &GetCommentByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByIDLogic) GetCommentByID(in *pb.GetCommentByIDReq) (*pb.GetCommentByIDResp, error) {
	findOneResp, err := l.svcCtx.CommentsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentByID find comment failed: %v", err)
	}

	return &pb.GetCommentByIDResp{
		Comment: &pb.Comment{
			Id:         findOneResp.Id,
			FeedID:     findOneResp.FeedId,
			UserID:     findOneResp.UserId,
			Content:    findOneResp.Content,
			Media0:     tool.GenMediaURL(findOneResp.Media0, l.svcCtx.Config.BaseURL),
			Media1:     tool.GenMediaURL(findOneResp.Media1, l.svcCtx.Config.BaseURL),
			Media2:     tool.GenMediaURL(findOneResp.Media2, l.svcCtx.Config.BaseURL),
			Media3:     tool.GenMediaURL(findOneResp.Media3, l.svcCtx.Config.BaseURL),
			CreateTime: timestamppb.New(findOneResp.CreateAt),
		},
	}, nil
}
