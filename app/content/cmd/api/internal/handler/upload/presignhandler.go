package upload

import (
	"net/http"

	"github.com/me2seeks/echo-hub/common/result"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/logic/upload"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PresignHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PresignReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upload.NewPresignLogic(r.Context(), svcCtx)
		resp, err := l.Presign(&req)
		result.HTTPResult(r, w, resp, err)
	}
}
