// Code generated by goctl. DO NOT EDIT.
package types

type GetContentCounterReq struct {
	ID        int64 `path:"id"`
	IsComment bool  `form:"is_comment"`
}

type GetContentCounterResp struct {
	CommentCount int64 `json:"comment_count"`
	LikeCount    int64 `json:"like_count"`
	ViewCount    int64 `json:"view_count"`
}

type GetUserCounterReq struct {
	UserID int64 `path:"id"`
}

type GetUserCounterResp struct {
	FollowingCount int64 `json:"following_count"`
	FollowerCount  int64 `json:"follower_count"`
	FeedCount      int64 `json:"feed_count"`
}