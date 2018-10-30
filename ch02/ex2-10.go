// Listing 2.10  The generateHTML function

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFile(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
