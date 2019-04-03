// Listing 6.16  Retrieving the relationship

func GetPost(id int) (post Post, err error) {
	post = Post{}
	  post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment(Post: &post)
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}
