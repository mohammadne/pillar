package config

import (
	"github.com/mohammadne/pillar/pkg/logger"
	"github.com/mohammadne/pillar/pkg/tracing"
)

type Config struct {
	Logger  *logger.Config `koanf:"logger"`
	Tracing tracing.Config `koanf:"tracing"`
}
