// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "github.com/me2seeks/echo-hub/app/content/cmd/api/internal/handler/comment"
	feed "github.com/me2seeks/echo-hub/app/content/cmd/api/internal/handler/feed"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// get comment comment list by page
				Method:  http.MethodGet,
				Path:    "/:commentID",
				Handler: comment.ListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/comment"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// create comment
				Method:  http.MethodPost,
				Path:    "/",
				Handler: comment.CreateHandler(serverCtx),
			},
			{
				// delete comment
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: comment.DeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/comment"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// get comment list by page
				Method:  http.MethodGet,
				Path:    "/:feedID/comment/",
				Handler: feed.ListCommentHandler(serverCtx),
			},
			{
				// get feed list by page
				Method:  http.MethodGet,
				Path:    "/:userID",
				Handler: feed.ListFeedHandler(serverCtx),
			},
		},
		rest.WithPrefix("/feed"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// create feed
				Method:  http.MethodPost,
				Path:    "/",
				Handler: feed.CreateHandler(serverCtx),
			},
			{
				// delete feed
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: feed.DeleteHandler(serverCtx),
			},
			{
				// create feed comment
				Method:  http.MethodPost,
				Path:    "/:feedID/comment/",
				Handler: feed.CreateCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/feed"),
	)
}