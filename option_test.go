package fun

import "testing"

func TestApplyNOptions(t *testing.T) {
	type Config struct {
		Port int
		Host string
	}

	var cfg Config
	ApplyNOptions(
		&cfg,
		func(c *Config) { c.Port = 8080 },
		func(c *Config) { c.Host = "localhost" },
	)
	if cfg.Port != 8080 {
		t.Errorf("expected port to be 8080, got %d", cfg.Port)
	}

	if cfg.Host != "localhost" {
		t.Errorf("expected host to be localhost, got %s", cfg.Host)
	}
}
