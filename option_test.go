package fun

import "testing"

func TestApplyOptions(t *testing.T) {
	type Config struct {
		Port int
		Host string
	}

	var cfg Config
	opts := []Option[Config]{
		func(c *Config) { c.Port = 8080 },
		func(c *Config) { c.Host = "localhost" },
	}

	ApplyOptions(&cfg, opts)
	if cfg.Port != 8080 {
		t.Errorf("expected port to be 8080, got %d", cfg.Port)
	}

	if cfg.Host != "localhost" {
		t.Errorf("expected host to be localhost, got %s", cfg.Host)
	}
}
