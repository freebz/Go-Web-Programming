// Listing 2.14  The Db global variable and the init function in data.go

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
