package mqs

type User struct {
	ID       int64
	Nickname string
	Handle   string
}

type Feed struct {
	ID      int64
	UserID  int64
	Content string
}

type SearchUsersResponse struct {
	Hits struct {
		Hits []struct {
			Source User `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type SearchFeedsResponse struct {
	Hits struct {
		Hits []struct {
			Source Feed `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
