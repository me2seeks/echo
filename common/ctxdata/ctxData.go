package ctxdata

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserID = "jwtUserId"

func GetUIDFromCtx(ctx context.Context) int64 {
	var uID int64
	uIDStr, ok := ctx.Value(CtxKeyJwtUserID).(string)
	if ok {
		logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", uID)
	}
	uID, err := strconv.ParseInt(uIDStr, 10, 64)
	if err != nil {
		logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
	}
	return uID
}
