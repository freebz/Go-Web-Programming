// Listing 6.8  Creating a post

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id "
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return
	}
	return
}
