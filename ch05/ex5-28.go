// Listing 5.28  A handler using the same template in different template files

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}
