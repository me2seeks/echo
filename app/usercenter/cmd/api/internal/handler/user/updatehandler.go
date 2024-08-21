package user

import (
	"net/http"

	"github.com/me2seeks/echo-hub/common/result"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/logic/user"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// update user info
func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		result.HTTPResult(r, w, resp, err)
	}
}
