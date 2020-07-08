package mt

type Config struct {
	Seed uint64
}

type Option func(*Config)

func WithSeed(seed uint64) Option {
	return func(c *Config) {
		c.Seed = seed
	}
}
