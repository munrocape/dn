package dnclient

type Stories struct {
	Stories []Story
}

type Story struct {
	Comment string `json:"comment"`
	Comments []string `json:"comments"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	SiteUrl string `json:"site_url"`
	Title string `json:"title"`
	Url string `json:"url"`
	VoteCount int `json:"vote_count"`
	Sponsored bool `json:"sponsored"`
}
