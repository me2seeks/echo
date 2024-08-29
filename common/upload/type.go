package upload

type ObjectType int32

const (
	Avatar ObjectType = iota
	FeedImg
	FeedVideo
	FeedGIF
)
