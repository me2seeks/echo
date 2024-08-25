package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数错误"
	message[TokenExpireError] = "token失效,请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[EncryptError] = "加密失败"
	message[FollowError] = "关注失败"
	message[UnFollowError] = "取消关注失败"
	message[CopyError] = "拷贝失败"
	message[KqPusherError] = "推送失败"
	message[MarshalError] = "序列化失败"
	message[UnmarshalError] = "反序列化失败"
	message[InvalidEvent] = "无效事件"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	}
	return "服务器开小差啦,稍后再来试一试"
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	}
	return false
}
