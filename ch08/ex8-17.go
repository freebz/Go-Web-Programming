// Listing 8.17  The new fetch method

func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
