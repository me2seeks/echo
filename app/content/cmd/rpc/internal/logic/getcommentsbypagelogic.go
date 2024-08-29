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

type GetCommentsByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByPageLogic {
	return &GetCommentsByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsByPageLogic) GetCommentsByPage(in *pb.GetCommentsByPageReq) (*pb.GetCommentsByPageResp, error) {
	var comments []*model.Comments
	var total int64
	var err error
	if !in.IsComment {
		comments, total, err = l.svcCtx.CommentsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().
			// Columns("id, user_id, content, media0, media1, media2, media3, create_at").
			Where("feed_id = ?", in.Id), in.Page, in.PageSize, "id DESC")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentListByPage FindPageListByPageWithTotal feedID%d err:%v", in.Id, err)
		}
	} else {
		comments, total, err = l.svcCtx.CommentsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().
			// Columns("id, user_id, content, media0, media1, media2, media3, create_at").
			Where("parent_id = ?", in.Id), in.Page, in.PageSize, "id DESC")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentListByPage FindPageListByPageWithTotal commentID%d err:%v", in.Id, err)
		}
	}

	resp := &pb.GetCommentsByPageResp{}
	resp.Total = total

	for _, comment := range comments {
		resp.Comments = append(resp.Comments, &pb.Comment{
			Id:         comment.Id,
			UserID:     comment.UserId,
			Content:    comment.Content,
			Media0:     tool.GenMediaURL(comment.Media0, l.svcCtx.Config.BaseURL),
			Media1:     tool.GenMediaURL(comment.Media1, l.svcCtx.Config.BaseURL),
			Media2:     tool.GenMediaURL(comment.Media2, l.svcCtx.Config.BaseURL),
			Media3:     tool.GenMediaURL(comment.Media3, l.svcCtx.Config.BaseURL),
			CreateTime: timestamppb.New(comment.CreateAt),
		})
	}

	return resp, nil
}
