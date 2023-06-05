package config

import (
	"github.com/mohammadne/pillar/pkg/logger"
	"github.com/mohammadne/pillar/pkg/tracing"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
		Tracing: tracing.Config{
			Enabled: false,
			Ratio:   1.0,
			Agent: tracing.Agent{
				Host: "127.0.0.1",
				Port: "6831",
			},
		},
	}
}
