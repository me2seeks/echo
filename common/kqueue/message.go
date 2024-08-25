// KqMessage
package kqueue

type EventType int

const (
	Follow EventType = iota
	UnFollow
	Like
	UnLike
	Comment
	UnComment
	Repost
	UnRepost
	View
	Feed
	DeleteFeed
)

type Event struct {
	Type      EventType
	ID        int64
	IsComment bool
}
