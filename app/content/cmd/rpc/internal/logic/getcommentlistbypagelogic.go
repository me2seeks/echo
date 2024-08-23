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

type GetCommentListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListByPageLogic {
	return &GetCommentListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListByPageLogic) GetCommentListByPage(in *pb.GetCommentListByPageReq) (*pb.GetCommentListByPageResp, error) {
	comments, total, err := l.svcCtx.CommentsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().Columns("*").Where("feed_id = ?", in.FeedID), in.Page, in.PageSize, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentListByPage FindPageListByPageWithTotal feedID%d err:%v", in.FeedID, err)
	}

	var commentList []*pb.Comment
	err = copier.Copy(&commentList, &comments)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CopyError), "GetCommentListByPage copier.Copy err:%v", err)
	}

	return &pb.GetCommentListByPageResp{
		Total:   total,
		Comment: commentList,
	}, nil
}
