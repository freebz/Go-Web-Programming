// Listing 4.7  Retrieving uploaded files using FormFile

func process(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.Readall(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}
