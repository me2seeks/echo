package es

import "time"

type User struct {
	ID        int64
	Handle    string
	Nickname  string
	Avatar    string
	CreatedAt time.Time
}

type Feed struct {
	ID        int64
	UserID    int64
	Content   string
	CreatedAt time.Time
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type UsersHit struct {
	Index  string  `json:"_index"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source User    `json:"_source"`
}

type UsersHits struct {
	Total    Total      `json:"total"`
	MaxScore float64    `json:"max_score"`
	Hits     []UsersHit `json:"hits"`
}

type SearchUsersResponse struct {
	Took     int       `json:"took"`
	TimedOut bool      `json:"timed_out"`
	Shards   Shards    `json:"_shards"`
	Hits     UsersHits `json:"hits"`
}

type FeedsHit struct {
	Index  string  `json:"_index"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Feed    `json:"_source"`
}

type FeedsHits struct {
	Total    Total      `json:"total"`
	MaxScore float64    `json:"max_score"`
	Hits     []FeedsHit `json:"hits"`
}

type SearchFeedsResponse struct {
	Took     int       `json:"took"`
	TimedOut bool      `json:"timed_out"`
	Shards   Shards    `json:"_shards"`
	Hits     FeedsHits `json:"hits"`
}
