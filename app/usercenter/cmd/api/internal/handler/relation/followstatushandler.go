package relation

import (
	"net/http"

	"github.com/me2seeks/echo-hub/common/result"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/logic/relation"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get follow status
func FollowStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := relation.NewFollowStatusLogic(r.Context(), svcCtx)
		resp, err := l.FollowStatus(&req)
		result.HTTPResult(r, w, resp, err)
	}
}
