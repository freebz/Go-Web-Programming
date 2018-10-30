// Listing 2.2  The index handler function in main.go

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		              "templates/navbar.html",
		              "templates/index.html",}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads(); if err == nil {
		templetes.ExecuteTemplate(w, "layout", threads)
	}
}
