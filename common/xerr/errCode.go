package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const (
	ServerCommonError uint32 = 100001 + iota
	RequestParamError
	TokenExpireError
	TokenGenerateError
	DbError
	DbUpdateAffectedZeroError
	EncryptError
	CopyError
	KqPusherError
	MarshalError
	UnmarshalError
	InvalidEvent
	EsError
	Str2Int64Error
	ReadBodyError
)

// 用户模块

// relation
const (
	FollowError uint32 = 200001 + iota
	UnFollowError
)
