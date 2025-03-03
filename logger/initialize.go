package logger

import (
	"go.uber.org/zap"
)

// Options is used to implement exact implementation for logger
type Options interface {
	GetOptions(production bool) (cfg zap.Config)
}

// InitLogger creates a default instance of zap logger
func InitLogger(production bool, options Options) (*zap.Logger, error) {
	cfg := options.GetOptions(production)
	return cfg.Build()
}
