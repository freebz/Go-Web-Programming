// Listing 4.1  The URL struct

type URL struct {
	Scheme   string
	Opaque   string
	User     *Userinfo
	Host     string
	Path     string
	RawQuery string
	Fragment string
}
