// Listing 9.19  The DB struct

type DB struct {
	mutex *sync.Mutex
	store map[string] [3]float64
}
