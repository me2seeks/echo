package ctxdata

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserID = "jwtUserId"

func GetUIDFromCtx(ctx context.Context) uint64 {
	var uid uint64
	if jsonUID, ok := ctx.Value(CtxKeyJwtUserID).(json.Number); ok {
		if int64UID, err := jsonUID.Int64(); err == nil {
			if int64UID < 0 {
				logx.WithContext(ctx).Errorf("Negative value cannot be converted to uint64  %d", jsonUID)
			}
			uid = uint64(int64UID)
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
