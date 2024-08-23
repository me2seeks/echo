//nolint:dupl
package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// comment
func (l *GetCommentListLogic) GetCommentList(in *pb.GetCommentListReq) (*pb.GetCommentListResp, error) {
	comments, err := l.svcCtx.CommentsModel.FindAll(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().Columns("*").Where("feed_id=?", in.FeedID), "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentList FindAll feedID%d err:%v", in.FeedID, err)
	}

	var commentList []*pb.Comment
	err = copier.Copy(&commentList, &comments)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CopyError), "GetCommentList copier.Copy err:%v", err)
	}

	return &pb.GetCommentListResp{}, nil
}
