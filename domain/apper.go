package domain

import (
	"alcoscanxlsx/config"

	"go.uber.org/zap"
)

type Apper interface {
	Options() *config.Configuration
	SaveOptions(key string, value interface{}) error
	SaveAllOptions() error
	Logger() *zap.SugaredLogger
	ConfigPath() string
	DbPath() string
	LogPath() string
}
