// Listing 2.15  The Threads function in thread.go

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread()
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId,
			&th.CreatedAt); err != nil {
				return
			}
		threads = append(threads, th)
	}
	rows.Close()
	return
}
