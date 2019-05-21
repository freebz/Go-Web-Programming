// Listing 8.15  The modified main function

func main() {

	var err error
	db , err := sql.Open("postgres", "user=gwp dbname=gwp password=gwp ssl-mode=disable")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post/", hndleRequest(&Post{Db: db}))
	server.ListenAndServe()
}
