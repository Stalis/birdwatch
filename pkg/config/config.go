package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Stalis/birdwatch/pkg/log"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/providers/structs"
	"github.com/spf13/pflag"
)

//
const ErrNoSuchFileOrDirectory = "open .*: no such file or directory"

var config *Config

const (
	LoggingLevelError = "Error"
)

// Server configuration.
type Config struct {
	// Port with api server is listen
	Port int `koanf:"port"`
	// Host name, using for listening
	Host string `koanf:"host"`
	// Logger configuration
	Logging LogConfig `koanf:"logging"`
	// Memory wather configuration
	Memory MemoryWatcherConfig `koanf:"memory"`
}

// Logger configuration.
type LogConfig struct {
	// Restrict logger print output to console
	Verbose bool `koanf:"verbose"`
	// Logging level
	Level string `koanf:"level"`
	// File to logging output
	File string `koanf:"file"`
}

// Memory watcher configuration.
type MemoryWatcherConfig struct {
	// If false - disabling memory watcher
	Enabled bool `koanf:"enabled"`
	// Interval for scanning memory state
	ScanInterval time.Duration `koanf:"scan_interval"`
}

// Return configuration struct, if initialized,
// else - try to initialize and return configuration struct
// or error if it fails.
func Get() (*Config, error) {
	if config == nil {
		res, err := InitConfig()
		if err != nil {
			return nil, err
		}
		config = res
	}

	return config, nil
}

// Try to initialize configuration struct
func InitConfig() (*Config, error) {
	f := initFlagSet()

	k := koanf.New(".")
	k, err := loadDefaultValues(k)
	if err != nil {
		return nil, err
	}

	k, err = initConfiguration(k, f)
	if err != nil {
		return nil, err
	}

	res := &Config{}

	if err := k.Unmarshal("", res); err != nil {
		return nil, err
	}

	return res, nil
}

func initFlagSet() *pflag.FlagSet {
	f := pflag.NewFlagSet("config", pflag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.IntP("port", "p", 50051, "Port for GRPC server")
	f.StringP("host", "h", "localhost", "Host for GRPC server")
	f.StringP("config", "c", "config.yaml", "config.yaml file path")
	f.String("logging-level", log.ErrorLevel, fmt.Sprintf("Logging level(one of %v)", log.LevelsList()))
	f.String("logging-file", "server.log", "Logging file")
	f.BoolP("logging-verbose", "v", false, "Verbose mode(logging to stdout/stderr)")

	f.SetNormalizeFunc(wordSeparationNormalizeFunc)

	f.Parse(os.Args[1:])
	return f
}

func loadDefaultValues(k *koanf.Koanf) (*koanf.Koanf, error) {
	err := k.Load(structs.Provider(Config{
		Port: 50051,
		Host: "localhost",
		Memory: MemoryWatcherConfig{
			Enabled:      true,
			ScanInterval: time.Second,
		},
		Logging: LogConfig{
			Verbose: false,
			Level:   log.ErrorLevel,
			File:    "server.log",
		},
	}, ""), nil)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func initConfiguration(k *koanf.Koanf, f *pflag.FlagSet) (*koanf.Koanf, error) {
	confPath, err := f.GetString("config")
	if err != nil {
		return nil, err
	}

	if err := k.Load(file.Provider(confPath), yaml.Parser()); err != nil {
		if match, _ := regexp.MatchString(ErrNoSuchFileOrDirectory, err.Error()); !match {
			return nil, err
		}
	}

	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return nil, err
	}

	return k, nil
}

func wordSeparationNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.ReplaceAll(name, sep, to)
	}
	return pflag.NormalizedName(name)
}
