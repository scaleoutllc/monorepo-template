package loglevel

import (
	"os"

	"github.com/labstack/gommon/log"
)

var (
	defaultLogLevel = log.INFO

	supportedLevels = map[string]log.Lvl{
		"DEBUG": log.DEBUG,
		"INFO":  log.INFO,
		"WARN":  log.WARN,
		"ERROR": log.ERROR,
	}
)

func GetLogLevel() log.Lvl {
	level, ok := supportedLevels[os.Getenv("APP_LOG_LEVEL")]

	if !ok {
		return defaultLogLevel
	}

	return level
}
