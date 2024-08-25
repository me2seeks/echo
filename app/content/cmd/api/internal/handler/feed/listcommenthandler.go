package feed

import (
	"net/http"

	"github.com/me2seeks/echo-hub/common/result"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/logic/feed"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get comment list by page
func ListCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFeedCommentsByPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := feed.NewListCommentLogic(r.Context(), svcCtx)
		resp, err := l.ListComment(&req)
		result.HTTPResult(r, w, resp, err)
	}
}
