package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	stdOut = "stdout"
	stdErr = "stderr"
)

type (
	// StdOptions is used for default console logs display
	StdOptions zap.Config
	// FileOptions is used for console logging and saving logs to file
	FileOptions zap.Config
)

var defaultCfg = zap.Config{
	Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
	Encoding: "json",
	EncoderConfig: zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	},
}

func (opt StdOptions) GetOptions(production bool) zap.Config {
	cfg := defaultCfg

	cfg.Development = !production
	cfg.OutputPaths = append(cfg.OutputPaths, stdOut)
	cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, stdErr)

	return cfg
}

// GetOptions is an extended configuration to log and save stdout & stderr lines in file
func (opt FileOptions) GetOptions(production bool, appLogsPath string, errLogsPath string) zap.Config {
	cfg := defaultCfg

	if appLogsPath == "" {
		appLogsPath = "logs/app.log"
	}

	if errLogsPath == "" {
		errLogsPath = "logs/error.log"
	}

	cfg.Development = !production
	cfg.OutputPaths = append(cfg.OutputPaths, stdOut, appLogsPath)
	cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, stdErr, errLogsPath)

	return cfg
}
