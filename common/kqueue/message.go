// KqMessage
package kqueue

type EventType int

const (
	// count
	Follow EventType = iota
	UnFollow
	Like
	UnLike
	Comment
	UnComment
	Repost
	UnRepost
	View

	// es
	Register

	// 共用
	Feed
	DeleteFeed
)

type CountEvent struct {
	Type      EventType
	ID        int64
	IsComment bool
}

type EsEvent struct {
	Type EventType
	// TODO 添加commentID
	ID       int64 // userID  feedID
	Nickname string
	Content  string
}
