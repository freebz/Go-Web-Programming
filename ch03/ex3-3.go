// Listing 3.3  The Server struct configuration

type Server struct {
	Addr           string
	Handler        Handler
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	TLSConfig      *tls.Config
	TLSNextProto   map[string]func(*Server, *tls.Conn, Handler)
	ConnState      func(net.Conn, ConnState)
	ErrorLog       *logLogger
}
