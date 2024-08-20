package ctxdata

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserID = "jwtUserId"

func GetUIDFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserID).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
