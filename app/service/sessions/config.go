package sessions

// Config provides the session configuration.
type Config struct {
	Name        string
	Secret      string
	MappingFile string
	Expiration  int64 // Expiration in seconds
	Inactivity  int64 // Inactivity in seconds
	Domain      string
	Secure      bool
}
