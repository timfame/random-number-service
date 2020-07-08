package grpc

type Config struct {
	host string
	port uint32
	seed uint64
}

type Option func(*Config)

func WithHost(host string) Option {
	return func(c *Config) {
		c.host = host
	}
}

func WithPort(port uint32) Option {
	return func(c *Config) {
		c.port = port
	}
}

func WithSeed(seed uint64) Option {
	return func(c *Config) {
		c.seed = seed
	}
}
