// Listing 6.9  Retrieving a post

func GetPost(id int) (post Post, err error) {
	post = Post()
	err = Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
