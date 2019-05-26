// Listing 9.15  The cloneTilesDB function

var TILESDB map[string] [3]float64

func cloneTilesDB() map[string] [3]float64 {
	db := make(map[stirng] [3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}
