// KqMessage
package kqueue

import "time"

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
	SourceID  int64
	TargetID  int64
	IsComment bool
}

type EsEvent struct {
	Type      EventType
	ID        int64
	UserID    int64
	Handle    string
	Avatar    string
	Content   string
	CreatedAt time.Time
}
