package handler

import (
	"net/http"

	"github.com/me2seeks/echo-hub/common/result"

	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/logic"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get content counter
func getContentCounterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetContentCounterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetContentCounterLogic(r.Context(), svcCtx)
		resp, err := l.GetContentCounter(&req)
		result.HTTPResult(r, w, resp, err)
	}
}