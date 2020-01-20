package post

import "time"

type Post struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created"`
	Name    string    `json:"name"`
	Slug    string    `json:"slug"`
	Content string    `json:"content,omitempty"`
}

func (p *Post) GetYear() string {
	return p.Created.Format("2006")
}
