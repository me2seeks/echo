package user

import (
	"net/http"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/logic/user"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get user info
func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		result.HTTPResult(r, w, resp, err)
	}
}
