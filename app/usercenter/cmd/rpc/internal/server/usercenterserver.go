// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package server

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/logic"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UsercenterServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UsercenterServer) GetUserAuthByAuthKey(ctx context.Context, in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	l := logic.NewGetUserAuthByAuthKeyLogic(ctx, s.svcCtx)
	return l.GetUserAuthByAuthKey(in)
}

func (s *UsercenterServer) GetUserAuthByUserID(ctx context.Context, in *pb.GetUserAuthByUserIDReq) (*pb.GetUserAuthyUserIDResp, error) {
	l := logic.NewGetUserAuthByUserIDLogic(ctx, s.svcCtx)
	return l.GetUserAuthByUserID(in)
}

func (s *UsercenterServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *UsercenterServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

func (s *UsercenterServer) Follow(ctx context.Context, in *pb.FollowReq) (*pb.FollowResp, error) {
	l := logic.NewFollowLogic(ctx, s.svcCtx)
	return l.Follow(in)
}

func (s *UsercenterServer) Unfollow(ctx context.Context, in *pb.UnfollowReq) (*pb.UnfollowResp, error) {
	l := logic.NewUnfollowLogic(ctx, s.svcCtx)
	return l.Unfollow(in)
}

func (s *UsercenterServer) GetFollowers(ctx context.Context, in *pb.GetFollowersReq) (*pb.GetFollowersResp, error) {
	l := logic.NewGetFollowersLogic(ctx, s.svcCtx)
	return l.GetFollowers(in)
}

func (s *UsercenterServer) GetFollowings(ctx context.Context, in *pb.GetFollowingsReq) (*pb.GetFollowingsResp, error) {
	l := logic.NewGetFollowingsLogic(ctx, s.svcCtx)
	return l.GetFollowings(in)
}

func (s *UsercenterServer) GetFollowingeCount(ctx context.Context, in *pb.GetFollowingeCountReq) (*pb.GetFollowingeCountResp, error) {
	l := logic.NewGetFollowingeCountLogic(ctx, s.svcCtx)
	return l.GetFollowingeCount(in)
}

func (s *UsercenterServer) GetFollowerCount(ctx context.Context, in *pb.GetFollowerCountReq) (*pb.GetFollowerCountResp, error) {
	l := logic.NewGetFollowerCountLogic(ctx, s.svcCtx)
	return l.GetFollowerCount(in)
}

func (s *UsercenterServer) GetFollowStatus(ctx context.Context, in *pb.GetFollowStatusReq) (*pb.GetFollowStatusResp, error) {
	l := logic.NewGetFollowStatusLogic(ctx, s.svcCtx)
	return l.GetFollowStatus(in)
}

func (s *UsercenterServer) LastRequestTime(ctx context.Context, in *pb.LastRequestTimeReq) (*pb.LastRequestTimeResp, error) {
	l := logic.NewLastRequestTimeLogic(ctx, s.svcCtx)
	return l.LastRequestTime(in)
}
