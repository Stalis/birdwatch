package config

import (
	"time"

	"github.com/Stalis/birdwatch/pkg/log"
)

var defaultConfig = Config{
	Port: -1,
	Host: "0.0.0.0",
	Memory: MemoryWatcherConfig{
		Enabled:      true,
		ScanInterval: time.Second,
	},
	Logging: LogConfig{
		Verbose: false,
		Level:   log.ErrorLevel,
		File:    "./server.log",
	},
}
