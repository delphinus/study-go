package github

import (
	"time"
)

// IssuesURL means url for github API
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult means as so
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is a issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // markdown
}

// User is a user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
