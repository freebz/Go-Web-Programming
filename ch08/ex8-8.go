// Listing 8.8  Multiplexing handler and GET handler function

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func handleGet(w http.ResponseWriter, r * http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrivev(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")

	if err != nil {
		return
	}
	w.Header().SEt("Content-Type", "application/json")
	w.Write(output)
	return
}
