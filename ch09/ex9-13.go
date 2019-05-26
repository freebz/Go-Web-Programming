// Listing 9.13  The tilesDB function

func tilesDB() map[string] [3]float64 {
	fmt.Println("Start populating tiles db ...")
	db := make(map[string] [3]float64)
	files, _ := ioutil.REadDir("tiles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}
