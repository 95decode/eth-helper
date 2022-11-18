package chain

type Config struct {
	name         string
	id           int
	endpoint     string
	keystorePath string
}

func ParseChainConfig(cfg *Config) (*Config, error) {
	config := &Config{
		name:         cfg.name,
		id:           cfg.id,
		endpoint:     cfg.endpoint,
		keystorePath: cfg.keystorePath,
	}

	return config, nil
}
