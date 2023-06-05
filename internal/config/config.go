package config

import (
	"github.com/mohammadne/pillar/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
