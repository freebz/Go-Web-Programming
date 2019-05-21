// Listing 8.12  Interface to pass into handlePost

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}
