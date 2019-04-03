// Listing 6.15  Creating the relationship

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	)
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

