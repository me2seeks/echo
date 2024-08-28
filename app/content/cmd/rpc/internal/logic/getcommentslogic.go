package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsLogic) GetComments(in *pb.GetCommentsReq) (*pb.GetCommentsResp, error) {
	var comments []*model.Comments
	var err error
	if !in.IsComment {
		comments, err = l.svcCtx.CommentsModel.FindAll(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().
			// Columns("id, user_id, content, media0, media1, media2, media3, create_at").
			Where("feed_id = ?", in.Id), "id DESC")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentListByPage FindAll feedID%d err:%v", in.Id, err)
		}
	} else {
		comments, err = l.svcCtx.CommentsModel.FindAll(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().
			// Columns("id, user_id, content, media0, media1, media2, media3, create_at").
			Where("parent_id = ?", in.Id), "")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentListByPage FindAll commentID%d err:%v", in.Id, err)
		}
	}

	resp := &pb.GetCommentsResp{}

	for _, comment := range comments {
		resp.Comments = append(resp.Comments, &pb.Comment{
			Id:         comment.Id,
			UserID:     comment.UserId,
			Content:    comment.Content,
			Media0:     comment.Media0.String,
			Media1:     comment.Media1.String,
			Media2:     comment.Media2.String,
			Media3:     comment.Media3.String,
			CreateTime: timestamppb.New(comment.CreateAt),
		})
	}
	return resp, nil
}
