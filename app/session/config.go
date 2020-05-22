package session

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

// NewConfig returns a new session configuration.
func NewConfig(name, secret, mappingFile string, expiration, inactivity int64, secure bool) Config {
	return Config{
		Name:        name,
		Secret:      secret,
		MappingFile: mappingFile,
		Expiration:  expiration,
		Inactivity:  inactivity,
		Secure:      secure,
	}
}
