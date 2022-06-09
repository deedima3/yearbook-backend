package entity

type Blogpost struct {
	PostId   uint64 `db:"postID"`
	Content  string `db:"content"`
	Upvote   uint64 `db:"upvote"`
	Downvote uint64 `db:"downvote"`
	Pages    uint64 `db:"pages"`
	Title    string `db:"title"`
}

type BlogPosts []*Blogpost
