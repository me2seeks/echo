package comment

import (
	"net/http"

	"github.com/me2seeks/echo-hub/common/result"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/logic/comment"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// unlike
func UnlikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCommentLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewUnlikeLogic(r.Context(), svcCtx)
		resp, err := l.Unlike(&req)
		result.HTTPResult(r, w, resp, err)
	}
}
